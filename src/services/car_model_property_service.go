package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.UpdateCarModelPropertyRequest, dto.CreateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.UpdateCarModelPropertyRequest, dto.CreateCarModelPropertyRequest, dto.CarModelPropertyResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "Property.Category"}},
		},
	}
}

func (p *CarModelPropertyService) GetById(ctx context.Context, id int) (*dto.CarModelPropertyResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelPropertyService) Update(ctx context.Context, req *dto.UpdateCarModelPropertyRequest, id int) (*dto.CarModelPropertyResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelPropertyService) Create(ctx context.Context, req *dto.CreateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelPropertyService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
