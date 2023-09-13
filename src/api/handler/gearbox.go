package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		service: services.NewGearboxService(cfg),
	}
}

// CreateGearbox godoc
// @Summary Create a Gearbox
// @Description Create a Gearbox
// @Tags Gearbox
// @Accept json
// @produces json
// @Param Request body dto.CreateGearboxRequest true "Create a Gearbox"
// @Success 201 {object} helper.Response{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/gearbox/create [post]
// @Security AuthBearer
func (pch *GearboxHandler) CreateGearbox(ctx *gin.Context) {
	Create[dto.CreateGearboxRequest, dto.GearboxResponse](ctx, pch.service.Create)

}

// DeleteGearbox godoc
// @Summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearbox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/get/{id} [get]
// @Security AuthBearer
func (pch *GearboxHandler) GetGearbox(ctx *gin.Context) {
	Get[dto.GearboxResponse](ctx, pch.service.GetById)
}

// UpdateGearbox godoc
// @Summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearbox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} helper.Response{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/update/{id} [put]
// @Security AuthBearer
func (pch *GearboxHandler) UpdateGearbox(ctx *gin.Context) {
	Update[dto.UpdateGearboxRequest, dto.GearboxResponse](ctx, pch.service.Update)
}

// GetGearbox godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearbox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Gearbox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/delete/{id} [delete]
// @Security AuthBearer
func (pch *GearboxHandler) DeleteGearbox(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
