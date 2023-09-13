package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.UpdateCompanyRequest, dto.CreateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.UpdateCompanyRequest, dto.CreateCompanyRequest, dto.CompanyResponse]{
			Db:       db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []Preload{{Name: "Country"}},
		},
	}
}

func (p *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CompanyService) Update(ctx context.Context, req *dto.UpdateCompanyRequest, id int) (*dto.CompanyResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CompanyService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
