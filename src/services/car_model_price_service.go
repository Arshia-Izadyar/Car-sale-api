package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelPriceService struct {
	base *BaseService[models.CarModelPrice, dto.UpdateCarModelPriceRequest, dto.CreateCarModelPriceRequest, dto.CarModelPriceResponse]
}

func NewCarModelPriceService(cfg *config.Config) *CarModelPriceService {
	return &CarModelPriceService{
		base: &BaseService[models.CarModelPrice, dto.UpdateCarModelPriceRequest, dto.CreateCarModelPriceRequest, dto.CarModelPriceResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *CarModelPriceService) GetById(ctx context.Context, id int) (*dto.CarModelPriceResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelPriceService) Update(ctx context.Context, req *dto.UpdateCarModelPriceRequest, id int) (*dto.CarModelPriceResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelPriceService) Create(ctx context.Context, req *dto.CreateCarModelPriceRequest) (*dto.CarModelPriceResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelPriceService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
