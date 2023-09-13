package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/common"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"gorm.io/gorm"
)

type Preload struct {
	Name string
}

type BaseService[T, Tu, Tc, Tr any] struct {
	Db       *gorm.DB
	Logger   logging.Logger
	Preloads []Preload
}

func NewBaseService[T, Tu, Tc, Tr any](cfg *config.Config) *BaseService[T, Tu, Tc, Tr] {
	return &BaseService[T, Tu, Tc, Tr]{
		Db:     db.GetDB(),
		Logger: logging.NewLogger(cfg),
	}
}

func GetPreload(db *gorm.DB, preloads []Preload) *gorm.DB {
	for _, p := range preloads {
		err := db.Preload(p.Name).Error
		if err == nil {
			db = db.Preload(p.Name)
		} else {
			panic(err)
		}
	}
	return db
}

func (bs *BaseService[T, Tu, Tc, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	db := GetPreload(bs.Db, bs.Preloads)
	err := db.Model(&model).Where("id = ?", id).First(&model).Error
	if err != nil {
		bs.Logger.Error(err, logging.Postgres, logging.Get, "cant get property category", nil)
		return nil, err
	}
	return common.TypeConvert[Tr](model)
}

func (bs *BaseService[T, Tu, Tc, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap, err := common.TypeConvert[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	snakeMap["updated_at"] = &sql.NullTime{Valid: true, Time: time.Now()}
	snakeMap["updated_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}

	model := new(T)
	tx := bs.Db.WithContext(ctx).Begin()

	err = tx.Model(&model).Where("id = ?", id).Updates(snakeMap).Error
	if err != nil {
		tx.Rollback()
		bs.Logger.Error(err, logging.Postgres, logging.Update, "cant update model", nil)
		return nil, err
	}
	tx.Commit()
	return bs.GetById(ctx, id)
}

func (bs *BaseService[T, Tu, Tc, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, err := common.TypeConvert[T](req)
	if err != nil {
		return nil, err
	}
	tx := bs.Db.WithContext(ctx).Begin()
	err = tx.Create(&model).Error
	if err != nil {
		tx.Rollback()
		bs.Logger.Error(err, logging.Postgres, logging.Update, "cant create model", nil)
		return nil, err
	}
	tx.Commit()
	bm, _ := common.TypeConvert[models.BaseModel](model)
	return bs.GetById(ctx, bm.ID)
}

func (bs *BaseService[T, Tu, Tc, Tr]) Delete(ctx context.Context, id int) error {
	tx := bs.Db.WithContext(ctx).Begin()
	model := new(T)
	if err := tx.First(&model, id).Error; err != nil {
		// If the record doesn't exist, rollback the transaction and return the error
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		bs.Logger.Error(err, logging.Postgres, logging.Get, "can't delete property category", nil)
		return err
	}

	// Delete the record
	if err := tx.Delete(&model).Error; err != nil {
		// If an error occurs during deletion, rollback the transaction and return the error
		tx.Rollback()
		bs.Logger.Error(err, logging.Postgres, logging.Get, "can't delete property category", nil)
		return err
	}
	tx.Commit()
	return nil
}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEquals":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}

				}
			}
		}
	}
	return strings.Join(query, " AND ")
}
func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

func NewPageList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PageList[T] {
	pl := &dto.PageList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPervious = pl.PageNumber > 1
	return pl

}

func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, preloads []Preload, db *gorm.DB) (*dto.PageList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr
	db = GetPreload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0

	err := db.
		Model(model).
		Where(query).
		Count(&totalRows).
		Error

	if err != nil {
		return nil, err
	}
	err = db.
		Where(query).
		Offset(pagination.GetOffSet()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}
	rItems, err = common.TypeConvert[[]Tr](items)
	if err != nil {
		return nil, err
	}
	return NewPageList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}

func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[Tr], error) {
	return Paginate[T, Tr](req, s.Preloads, s.Db)

}
