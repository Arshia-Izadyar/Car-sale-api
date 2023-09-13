package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CarModelFileHandler struct {
	service *services.CarModelFileService
}

func NewCarModelFileHandler(cfg *config.Config) *CarModelFileHandler {
	return &CarModelFileHandler{
		service: services.NewCarModelFileService(cfg),
	}
}

// CreateCarModelFile godoc
// @Summary Create a CarModelFile
// @Description Create a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelFileRequest true "Create a CarModelFile"
// @Success 201 {object} helper.Response{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-file/create [post]
// @Security AuthBearer
func (pch *CarModelFileHandler) CreateCarModelFile(ctx *gin.Context) {
	Create[dto.CreateCarModelFileRequest, dto.CarModelFileResponse](ctx, pch.service.Create)

}

// DeleteCarModelFile godoc
// @Summary Delete a CarModelFile
// @Description Delete a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/get/{id} [get]
// @Security AuthBearer
func (pch *CarModelFileHandler) GetCarModelFile(ctx *gin.Context) {
	Get[dto.CarModelFileResponse](ctx, pch.service.GetById)
}

// UpdateCarModelFile godoc
// @Summary Update a CarModelFile
// @Description Update a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelFileRequest true "Update a CarModelFile"
// @Success 200 {object} helper.Response{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/update/{id} [put]
// @Security AuthBearer
func (pch *CarModelFileHandler) UpdateCarModelFile(ctx *gin.Context) {
	Update[dto.UpdateCarModelFileRequest, dto.CarModelFileResponse](ctx, pch.service.Update)
}

// GetCarModelFile godoc
// @Summary Get a CarModelFile
// @Description Get a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/delete/{id} [delete]
// @Security AuthBearer
func (pch *CarModelFileHandler) DeleteCarModelFile(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
