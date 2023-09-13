package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/gin-gonic/gin"
)

func Create[Ti, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	request := new(Ti)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := caller(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))
}

func Get[To any](ctx *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), "id value is not valid"))
	}
	res, err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, int(helper.Success), true))
}

func Update[Ti, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti, id int) (*To, error)) {
	req := new(Ti)
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), "id value is not valid"))
	}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
		return
	}
	res, err := caller(ctx, req, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, int(helper.Success), true))
}

func Delete(ctx *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), "id value is not valid"))
	}
	err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, int(helper.Success), true))
}
