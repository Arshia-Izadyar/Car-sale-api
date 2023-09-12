package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type GenericPropertyService struct {
	base *BaseService[models.Property, dto.UpdatePropertyRequest, dto.CreatePropertyRequest, dto.PropertyResponse]
}

func NewGenericPropertyService(cfg *config.Config) *GenericPropertyService {
	return &GenericPropertyService{
		base: &BaseService[models.Property, dto.UpdatePropertyRequest, dto.CreatePropertyRequest, dto.PropertyResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "Category"}},
		},
	}
}

func (p *GenericPropertyService) GetById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *GenericPropertyService) Update(ctx context.Context, req *dto.UpdatePropertyRequest, id int) (*dto.PropertyResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *GenericPropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *GenericPropertyService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
