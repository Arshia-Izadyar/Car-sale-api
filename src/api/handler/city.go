package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags City
// @Accept json
// @produces json
// @Param Request body dto.CreateCityRequest true "Create a City"
// @Success 201 {object} helper.Response{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/city/create [post]
// @Security AuthBearer
func (pch *CityHandler) CreateCity(ctx *gin.Context) {
	Create[dto.CreateCityRequest, dto.CityResponse](ctx, pch.service.Create)

}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags City
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/city/get/{id} [get]
// @Security AuthBearer
func (pch *CityHandler) GetCity(ctx *gin.Context) {
	Get[dto.CityResponse](ctx, pch.service.GetById)
}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags City
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCityRequest true "Update a City"
// @Success 200 {object} helper.Response{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/city/update/{id} [put]
// @Security AuthBearer
func (pch *CityHandler) UpdateCity(ctx *gin.Context) {
	Update[dto.UpdateCityRequest, dto.CityResponse](ctx, pch.service.Update)
}

// GetCity godoc
// @Summary Get a City
// @Description Get a City
// @Tags City
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/city/delete/{id} [delete]
// @Security AuthBearer
func (pch *CityHandler) DeleteCity(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
