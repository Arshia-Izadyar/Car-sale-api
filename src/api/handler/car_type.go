package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		service: services.NewCarTypeService(cfg),
	}
}

// CreateCarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} helper.Response{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-type/create [post]
// @Security AuthBearer
func (pch *CarTypeHandler) CreateCarType(ctx *gin.Context) {
	Create[dto.CreateCarTypeRequest, dto.CarTypeResponse](ctx, pch.service.Create)

}

// DeleteCarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/get/{id} [get]
// @Security AuthBearer
func (pch *CarTypeHandler) GetCarType(ctx *gin.Context) {
	Get[dto.CarTypeResponse](ctx, pch.service.GetById)
}

// UpdateCarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} helper.Response{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/update/{id} [put]
// @Security AuthBearer
func (pch *CarTypeHandler) UpdateCarType(ctx *gin.Context) {
	Update[dto.UpdateCarTypeRequest, dto.CarTypeResponse](ctx, pch.service.Update)
}

// GetCarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/delete/{id} [delete]
// @Security AuthBearer
func (pch *CarTypeHandler) DeleteCarType(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
