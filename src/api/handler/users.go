package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
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

// SendOtp godoc
// @Summary SendOtp
// @Description SendOtp
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.GetOtpRequest true "Create a GetOtpRequest"
// @Success 201 {object} helper.Response "SendOtp response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/send-otp [post]
func (uh *UserHandler) SendOtp(ctx *gin.Context) {
	req := dto.GetOtpRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	err = uh.service.SendOtp(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(gin.H{"Status": "Sent!"}, int(helper.Success), true))
}

// RegisterLoginByPhone godoc
// @Summary RegisterLoginByPhone
// @Description RegisterLoginByPhone
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RegisterLoginByPhone true "Create a RegisterLoginByPhone"
// @Success 201 {object} helper.Response "RegisterLoginByPhone response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/register/phone [post]
// @Router /v1/users/login/phone [post]
func (uh *UserHandler) RegisterLoginByPhone(ctx *gin.Context) {
	req := dto.RegisterLoginByPhone{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	tk, err := uh.service.RegisterLoginByPhone(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(tk, int(helper.Success), true))

}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description RegisterByUsername
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RegisterUserByUsername true "Create a RegisterByUsername"
// @Success 201 {object} helper.Response "RegisterByUsername response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/register/username [post]
func (us *UserHandler) RegisterByUsername(ctx *gin.Context) {
	req := dto.RegisterUserByUsername{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	err = us.service.RegisterByUsername(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(gin.H{"Status": "user created please login"}, int(helper.Success), true))

}

// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.LoginByUsername true "Create a LoginByUsername"
// @Success 201 {object} helper.Response "LoginByUsername response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/login/username [post]
func (us *UserHandler) LoginByUsername(ctx *gin.Context) {
	req := dto.LoginByUsername{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
		return
	}
	res, err := us.service.LoginByUsername(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))

}

// RefreshToken godoc
// @Summary RefreshToken
// @Description RefreshToken
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RefreshToken true "Create a RefreshToken"
// @Success 201 {object} helper.Response "RefreshToken response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/refresh [post]
func (t *UserHandler) RefreshToken(ctx *gin.Context) {
	req := dto.RefreshToken{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := t.service.Token.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))

}

func (t *UserHandler) Logout(ctx *gin.Context) {
	authToken := ctx.GetHeader(constants.AuthenTicationHeaderKey)
	token := strings.Split(authToken, " ")[1]
	err := services.AddToBlackList(token, t.service.Cfg.Jwt.AccessTokenExpireDuration*time.Minute)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("logged out successfully", int(helper.Success), true))
}
