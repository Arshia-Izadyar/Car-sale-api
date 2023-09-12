package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
)

type GenericPropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.UpdatePropertyCategoryRequest, dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewGenericPropertyCategoryService(cfg *config.Config) *GenericPropertyCategoryService {
	base := NewBaseService[models.PropertyCategory, dto.UpdatePropertyCategoryRequest, dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse](cfg)
	return &GenericPropertyCategoryService{
		base: base,
	}
}

func (p *GenericPropertyCategoryService) GetById(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *GenericPropertyCategoryService) Update(ctx context.Context, req *dto.UpdatePropertyCategoryRequest, id int) (*dto.PropertyCategoryResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *GenericPropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *GenericPropertyCategoryService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
