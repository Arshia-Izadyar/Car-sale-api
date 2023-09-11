package handler

import (
	"net/http"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	srv := services.NewUserService(cfg)
	return &UserHandler{
		service: srv,
	}
}

func (uh *UserHandler) SendOtp(ctx *gin.Context) {
	req := dto.GetOtpRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	err = uh.service.SendOtp(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(gin.H{"Status": "Sent!"}, int(helper.Success), true))
}

func (uh *UserHandler) RegisterLoginByPhone(ctx *gin.Context) {
	req := dto.RegisterLoginByPhone{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	tk, err := uh.service.RegisterLoginByPhone(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(tk, int(helper.Success), true))

}

func (us *UserHandler) RegisterByUsername(ctx *gin.Context) {
	req := dto.RegisterUserByUsername{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	err = us.service.RegisterByUsername(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(gin.H{"Status": "user created please login"}, int(helper.Success), true))

}

func (us *UserHandler) LoginByUsername(ctx *gin.Context) {
	req := dto.LoginByUsername{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	res, err := us.service.LoginByUsername(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))

}
