package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		service: services.NewCarModelCommentService(cfg),
	}
}

// CreateCarModelComment godoc
// @Summary Create a CarModelComment
// @Description Create a CarModelComment
// @Tags CarModelComment
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.Response{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-comment/create [post]
// @Security AuthBearer
func (pch *CarModelCommentHandler) CreateCarModelComment(ctx *gin.Context) {
	Create[dto.CreateCarModelCommentRequest, dto.CarModelCommentResponse](ctx, pch.service.Create)

}

// DeleteCarModelComment godoc
// @Summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComment
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-comment/get/{id} [get]
// @Security AuthBearer
func (pch *CarModelCommentHandler) GetCarModelComment(ctx *gin.Context) {
	Get[dto.CarModelCommentResponse](ctx, pch.service.GetById)
}

// UpdateCarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComment
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 200 {object} helper.Response{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-comment/update/{id} [put]
// @Security AuthBearer
func (pch *CarModelCommentHandler) UpdateCarModelComment(ctx *gin.Context) {
	Update[dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse](ctx, pch.service.Update)
}

// GetCarModelComment godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComment
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelComment response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-comment/delete/{id} [delete]
// @Security AuthBearer
func (pch *CarModelCommentHandler) DeleteCarModelComment(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
