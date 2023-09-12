package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

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
