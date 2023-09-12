package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"gorm.io/gorm"
)

type PropertyCategoryService struct {
	logger logging.Logger
	db     *gorm.DB
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	logger := logging.NewLogger(cfg)
	db := db.GetDB()
	return &PropertyCategoryService{
		db:     db,
		logger: logger,
	}
}

func (ps *PropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	propertyCategory := &models.PropertyCategory{
		Name: req.Name,
		Icon: req.Icon,
	}
	tx := ps.db.WithContext(ctx).Begin()
	err := tx.Create(&propertyCategory).Error
	if err != nil {
		ps.logger.Error(err, logging.Postgres, logging.Insert, "cant add property category", nil)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	res := &dto.PropertyCategoryResponse{
		Id:   propertyCategory.ID,
		Name: propertyCategory.Name,
		Icon: propertyCategory.Icon,
	}
	return res, nil
}

func (ps *PropertyCategoryService) Get(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	property := &models.PropertyCategory{}
	err := ps.db.Model(&models.PropertyCategory{}).Where("id = ? AND deleted_by IS NULL", id).First(&property).Error
	if err != nil {
		ps.logger.Error(err, logging.Postgres, logging.Get, "cant get property category", nil)
		return nil, err
	}
	res := &dto.PropertyCategoryResponse{
		Id:   id,
		Name: property.Name,
		Icon: property.Icon,
	}
	return res, nil
}

func (ps *PropertyCategoryService) Update(ctx context.Context, req *dto.UpdatePropertyCategoryRequest, id int) (*dto.PropertyCategoryResponse, error) {
	updateMap := map[string]interface{}{
		"name":       req.Name,
		"updated_at": sql.NullTime{Time: time.Now(), Valid: true},
		"updated_by": sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64))},
	}
	tx := ps.db.WithContext(ctx).Begin()
	err := tx.Model(&models.PropertyCategory{}).Where("id = ? AND deleted_by is null", id).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		ps.logger.Error(err, logging.Postgres, logging.Get, "cant update property category", nil)
		return nil, err
	}
	updatedProperty := &models.PropertyCategory{}
	err = tx.Model(&models.PropertyCategory{}).Where("id = ? AND deleted_by is null", id).First(&updatedProperty).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	res := &dto.PropertyCategoryResponse{
		Id:   id,
		Name: updatedProperty.Name,
		Icon: updatedProperty.Icon,
	}
	return res, nil

}

func (ps *PropertyCategoryService) Delete(ctx context.Context, id int) error {
	tx := ps.db.WithContext(ctx).Begin()
	category := models.PropertyCategory{}
	if err := tx.First(&category, id).Error; err != nil {
		// If the record doesn't exist, rollback the transaction and return the error
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		ps.logger.Error(err, logging.Postgres, logging.Get, "can't delete property category", nil)
		return err
	}

	// Delete the record
	if err := tx.Delete(&category).Error; err != nil {
		// If an error occurs during deletion, rollback the transaction and return the error
		tx.Rollback()
		ps.logger.Error(err, logging.Postgres, logging.Get, "can't delete property category", nil)
		return err
	}
	tx.Commit()
	return nil
}
