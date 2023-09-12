package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.GenericPropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		service: services.NewGenericPropertyService(cfg),
	}
}

// CreateProperty godoc
// @Summary Create a Property
// @Description Create a Property
// @Tags Property
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyRequest true "Create a Property"
// @Success 201 {object} helper.Response{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/create [post]
// @Security AuthBearer
func (pch *PropertyHandler) CreateProperty(ctx *gin.Context) {
	Create[dto.CreatePropertyRequest, dto.PropertyResponse](ctx, pch.service.Create)

}

// DeleteProperty godoc
// @Summary Delete a Property
// @Description Delete a Property
// @Tags Property
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property/get/{id} [get]
// @Security AuthBearer
func (pch *PropertyHandler) GetProperty(ctx *gin.Context) {
	Get[dto.PropertyResponse](ctx, pch.service.GetById)
}

// UpdateProperty godoc
// @Summary Update a Property
// @Description Update a Property
// @Tags Property
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} helper.Response{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property/update/{id} [put]
// @Security AuthBearer
func (pch *PropertyHandler) UpdateProperty(ctx *gin.Context) {
	Update[dto.UpdatePropertyRequest, dto.PropertyResponse](ctx, pch.service.Update)
}

// GetProperty godoc
// @Summary Get a Property
// @Description Get a Property
// @Tags Property
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property/delete/{id} [delete]
// @Security AuthBearer
func (pch *PropertyHandler) DeleteProperty(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
