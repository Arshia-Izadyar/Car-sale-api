package handler

import (
	"net/http"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	s services.TokenService
}

func NewTestHandler(cfg *config.Config) *TestHandler {
	return &TestHandler{
		s: *services.NewTokenService(cfg),
	}
}

func (t *TestHandler) RefreshToken(ctx *gin.Context) {
	req := dto.RefreshToken{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := t.s.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))

}
