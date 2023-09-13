package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.UpdateCarModelRequest, dto.CreateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.UpdateCarModelRequest, dto.CreateCarModelRequest, dto.CarModelResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
			Preloads: []Preload{
				{Name: "Company.Country"},
				{Name: "CarType"},
				{Name: "Gearbox"},
				{Name: "CarModelYears.PersianYear"},
				{Name: "CarModelYears.CarModelPrice"},
				{Name: "CarModelColors"},
				{Name: "CarModelFiles.File"},
				{Name: "CarModelProperties.Property.Category"},
				{Name: "CarModelComments.User"},
			},
		},
	}
}

func (p *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelService) Update(ctx context.Context, req *dto.UpdateCarModelRequest, id int) (*dto.CarModelResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
