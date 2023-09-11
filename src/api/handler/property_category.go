package handler

import (
	"net/http"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

func (pch *PropertyCategoryHandler) Create(ctx *gin.Context) {
	req := dto.CreatePropertyCategoryRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	res, err := pch.service.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))

}
