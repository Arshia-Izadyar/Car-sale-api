package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.UpdateGearboxRequest, dto.CreateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.UpdateGearboxRequest, dto.CreateGearboxRequest, dto.GearboxResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *GearboxService) Update(ctx context.Context, req *dto.UpdateGearboxRequest, id int) (*dto.GearboxResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *GearboxService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
