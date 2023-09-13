package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CarModelYearHandler struct {
	service *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		service: services.NewCarModelYearService(cfg),
	}
}

// CreateCarModelYear godoc
// @Summary Create a CarModelYear
// @Description Create a CarModelYear
// @Tags CarModelYear
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelYearRequest true "Create a CarModelYear"
// @Success 201 {object} helper.Response{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-year/create [post]
// @Security AuthBearer
func (pch *CarModelYearHandler) CreateCarModelYear(ctx *gin.Context) {
	Create[dto.CreateCarModelYearRequest, dto.CarModelYearResponse](ctx, pch.service.Create)

}

// DeleteCarModelYear godoc
// @Summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-year/get/{id} [get]
// @Security AuthBearer
func (pch *CarModelYearHandler) GetCarModelYear(ctx *gin.Context) {
	Get[dto.CarModelYearResponse](ctx, pch.service.GetById)
}

// UpdateCarModelYear godoc
// @Summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.Response{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-year/update/{id} [put]
// @Security AuthBearer
func (pch *CarModelYearHandler) UpdateCarModelYear(ctx *gin.Context) {
	Update[dto.UpdateCarModelYearRequest, dto.CarModelYearResponse](ctx, pch.service.Update)
}

// GetCarModelYear godoc
// @Summary Get a CarModelYear
// @Description Get a CarModelYear
// @Tags CarModelYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-year/delete/{id} [delete]
// @Security AuthBearer
func (pch *CarModelYearHandler) DeleteCarModelYear(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
