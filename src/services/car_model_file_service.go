package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelFileService struct {
	base *BaseService[models.CarModelFile, dto.UpdateCarModelFileRequest, dto.CreateCarModelFileRequest, dto.CarModelFileResponse]
}

func NewCarModelFileService(cfg *config.Config) *CarModelFileService {
	return &CarModelFileService{
		base: &BaseService[models.CarModelFile, dto.UpdateCarModelFileRequest, dto.CreateCarModelFileRequest, dto.CarModelFileResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *CarModelFileService) GetById(ctx context.Context, id int) (*dto.CarModelFileResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelFileService) Update(ctx context.Context, req *dto.UpdateCarModelFileRequest, id int) (*dto.CarModelFileResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelFileService) Create(ctx context.Context, req *dto.CreateCarModelFileRequest) (*dto.CarModelFileResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *CarModelFileService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
