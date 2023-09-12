package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.UpdatePersianYearRequest, dto.CreatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.UpdatePersianYearRequest, dto.CreatePersianYearRequest, dto.PersianYearResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *PersianYearService) GetById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *PersianYearService) Update(ctx context.Context, req *dto.UpdatePersianYearRequest, id int) (*dto.PersianYearResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *PersianYearService) Create(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *PersianYearService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
