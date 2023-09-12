package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.UpdateColorRequest, dto.CreateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.UpdateColorRequest, dto.CreateColorRequest, dto.ColorResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
			// Preloads: []Preload{{Name: ""}},
		},
	}
}

func (p *ColorService) GetById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *ColorService) Update(ctx context.Context, req *dto.UpdateColorRequest, id int) (*dto.ColorResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *ColorService) Create(ctx context.Context, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *ColorService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
