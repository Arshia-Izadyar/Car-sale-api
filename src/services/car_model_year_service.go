package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.UpdateCarModelYearRequest, dto.CreateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.UpdateCarModelYearRequest, dto.CreateCarModelYearRequest, dto.CarModelYearResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "PersianYear"}},
		},
	}
}

func (p *CarModelYearService) GetById(ctx context.Context, id int) (*dto.CarModelYearResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelYearService) Update(ctx context.Context, req *dto.UpdateCarModelYearRequest, id int) (*dto.CarModelYearResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelYearService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
