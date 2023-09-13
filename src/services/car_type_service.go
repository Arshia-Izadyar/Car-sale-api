package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.UpdateCarTypeRequest, dto.CreateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.UpdateCarTypeRequest, dto.CreateCarTypeRequest, dto.CarTypeResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *CarTypeService) GetById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarTypeService) Update(ctx context.Context, req *dto.UpdateCarTypeRequest, id int) (*dto.CarTypeResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarTypeService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
