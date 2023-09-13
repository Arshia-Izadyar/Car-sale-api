package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CarModelPriceHandler struct {
	service *services.CarModelPriceService
}

func NewCarModelPriceHandler(cfg *config.Config) *CarModelPriceHandler {
	return &CarModelPriceHandler{
		service: services.NewCarModelPriceService(cfg),
	}
}

// CreateCarModelPrice godoc
// @Summary Create a CarModelPrice
// @Description Create a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPriceRequest true "Create a CarModelPrice"
// @Success 201 {object} helper.Response{result=dto.CarModelPriceResponse} "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-price/create [post]
// @Security AuthBearer
func (pch *CarModelPriceHandler) CreateCarModelPrice(ctx *gin.Context) {
	Create[dto.CreateCarModelPriceRequest, dto.CarModelPriceResponse](ctx, pch.service.Create)

}

// DeleteCarModelPrice godoc
// @Summary Delete a CarModelPrice
// @Description Delete a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/get/{id} [get]
// @Security AuthBearer
func (pch *CarModelPriceHandler) GetCarModelPrice(ctx *gin.Context) {
	Get[dto.CarModelPriceResponse](ctx, pch.service.GetById)
}

// UpdateCarModelPrice godoc
// @Summary Update a CarModelPrice
// @Description Update a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceRequest true "Update a CarModelPrice"
// @Success 200 {object} helper.Response{result=dto.CarModelPriceResponse} "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/update/{id} [put]
// @Security AuthBearer
func (pch *CarModelPriceHandler) UpdateCarModelPrice(ctx *gin.Context) {
	Update[dto.UpdateCarModelPriceRequest, dto.CarModelPriceResponse](ctx, pch.service.Update)
}

// GetCarModelPrice godoc
// @Summary Get a CarModelPrice
// @Description Get a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/delete/{id} [delete]
// @Security AuthBearer
func (pch *CarModelPriceHandler) DeleteCarModelPrice(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
