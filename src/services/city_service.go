package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.UpdateCityRequest, dto.CreateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.UpdateCityRequest, dto.CreateCityRequest, dto.CityResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "Country"}},
		},
	}
}

func (p *CityService) GetById(ctx context.Context, id int) (*dto.CityResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CityService) Update(ctx context.Context, req *dto.UpdateCityRequest, id int) (*dto.CityResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CityService) Create(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CityService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
