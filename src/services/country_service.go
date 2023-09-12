package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CountryService struct {
	base *BaseService[models.Country, dto.UpdateCountryRequest, dto.CreateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.UpdateCountryRequest, dto.CreateCountryRequest, dto.CountryResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "Cities"}},
		},
	}
}

func (p *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CountryService) Update(ctx context.Context, req *dto.UpdateCountryRequest, id int) (*dto.CountryResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CountryService) Create(ctx context.Context, req *dto.CreateCountryRequest) (*dto.CountryResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CountryService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
