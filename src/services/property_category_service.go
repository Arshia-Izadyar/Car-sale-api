package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
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
