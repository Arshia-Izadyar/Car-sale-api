package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.UpdateCarModelColorRequest, dto.CreateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.UpdateCarModelColorRequest, dto.CreateCarModelColorRequest, dto.CarModelColorResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *CarModelColorService) GetById(ctx context.Context, id int) (*dto.CarModelColorResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelColorService) Update(ctx context.Context, req *dto.UpdateCarModelColorRequest, id int) (*dto.CarModelColorResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelColorService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
