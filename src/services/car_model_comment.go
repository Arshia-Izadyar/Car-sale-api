package services

import (
	"context"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.UpdateCarModelCommentRequest, dto.CreateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.UpdateCarModelCommentRequest, dto.CreateCarModelCommentRequest, dto.CarModelCommentResponse]{
			Db:     db.GetDB(),
			Logger: logging.NewLogger(cfg),
			Preloads: []Preload{
				{Name: "User"},
			},
		},
	}
}

func (p *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return p.base.GetById(ctx, id)
}

func (p *CarModelCommentService) Update(ctx context.Context, req *dto.UpdateCarModelCommentRequest, id int) (*dto.CarModelCommentResponse, error) {
	return p.base.Update(ctx, req, id)
}

func (p *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return p.base.Create(ctx, req)
}

func (p *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return p.base.Delete(ctx, id)
}
