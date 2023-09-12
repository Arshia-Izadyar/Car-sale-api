package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type FileService struct {
	base *BaseService[models.File, dto.UpdateFileRequest, dto.CreateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.UpdateFileRequest, dto.CreateFileRequest, dto.FileResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
		},
	}
}

func (p *FileService) GetById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *FileService) Update(ctx context.Context, req *dto.UpdateFileRequest, id int) (*dto.FileResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *FileService) Create(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {
	return p.base.Create(ctx, req)
}

func (p *FileService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
