package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		service: services.NewColorService(cfg),
	}
}

// CreateColor godoc
// @Summary Create a Color
// @Description Create a Color
// @Tags Color
// @Accept json
// @produces json
// @Param Request body dto.CreateColorRequest true "Create a Color"
// @Success 201 {object} helper.Response{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/color/create [post]
// @Security AuthBearer
func (pch *ColorHandler) CreateColor(ctx *gin.Context) {
	Create[dto.CreateColorRequest, dto.ColorResponse](ctx, pch.service.Create)

}

// DeleteColor godoc
// @Summary Delete a Color
// @Description Delete a Color
// @Tags Color
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/color/get/{id} [get]
// @Security AuthBearer
func (pch *ColorHandler) GetColor(ctx *gin.Context) {
	Get[dto.ColorResponse](ctx, pch.service.GetById)
}

// UpdateColor godoc
// @Summary Update a Color
// @Description Update a Color
// @Tags Color
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} helper.Response{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/color/update/{id} [put]
// @Security AuthBearer
func (pch *ColorHandler) UpdateColor(ctx *gin.Context) {
	Update[dto.UpdateColorRequest, dto.ColorResponse](ctx, pch.service.Update)
}

// GetColor godoc
// @Summary Get a Color
// @Description Get a Color
// @Tags Color
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Color response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/color/delete/{id} [delete]
// @Security AuthBearer
func (pch *ColorHandler) DeleteColor(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
