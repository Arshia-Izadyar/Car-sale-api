package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a Country
// @Description Create a Country
// @Tags Country
// @Accept json
// @produces json
// @Param Request body dto.CreateCountryRequest true "Create a Country"
// @Success 201 {object} helper.Response{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/create [post]
// @Security AuthBearer
func (pch *CountryHandler) CreateCountry(ctx *gin.Context) {
	Create[dto.CreateCountryRequest, dto.CountryResponse](ctx, pch.service.Create)

}

// DeleteCountry godoc
// @Summary Delete a Country
// @Description Delete a Country
// @Tags Country
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/country/get/{id} [get]
// @Security AuthBearer
func (pch *CountryHandler) GetCountry(ctx *gin.Context) {
	Get[dto.CountryResponse](ctx, pch.service.GetById)
}

// UpdateCountry godoc
// @Summary Update a Country
// @Description Update a Country
// @Tags Country
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCountryRequest true "Update a Country"
// @Success 200 {object} helper.Response{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/country/update/{id} [put]
// @Security AuthBearer
func (pch *CountryHandler) UpdateCountry(ctx *gin.Context) {
	Update[dto.UpdateCountryRequest, dto.CountryResponse](ctx, pch.service.Update)
}

// GetCountry godoc
// @Summary Get a Country
// @Description Get a Country
// @Tags Country
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/country/delete/{id} [delete]
// @Security AuthBearer
func (pch *CountryHandler) DeleteCountry(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
