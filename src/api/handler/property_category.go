package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.GenericPropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewGenericPropertyCategoryService(cfg),
	}
}

// CreatePropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create a PropertyCategory"
// @Success 201 {object} helper.Response{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/create [post]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) CreatePropertyCategory(ctx *gin.Context) {
	// req := dto.CreatePropertyCategoryRequest{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.InternalError), err))
	// 	return
	// }
	// res, err := pch.service.Create(ctx, &req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, int(helper.Success), true))
	Create[dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, pch.service.Create)

}

// DeletePropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property-category/get/{id} [get]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) GetPropertyCategory(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id == 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("index 0 not found").Error()))
	// 	return
	// }
	// res, err := pch.service.GetById(ctx, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, int(helper.Success), true))
	Get[dto.PropertyCategoryResponse](ctx, pch.service.GetById)
}

// UpdatePropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 200 {object} helper.Response{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property-category/update/{id} [put]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) UpdatePropertyCategory(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id == 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("index 0 not found").Error()))
	// 	return
	// }
	// req := dto.UpdatePropertyCategoryRequest{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
	// 	return
	// }
	// res, err := pch.service.Update(ctx, &req, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, int(helper.Success), true))
	Update[dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, pch.service.Update)
}

// GetPropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a Property
// @Tags PropertyCategory
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "PropertyCategory response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/property-category/delete/{id} [delete]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) DeletePropertyCategory(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id == 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("index 0 not found").Error()))
	// 	return
	// }
	// err := pch.service.Delete(ctx, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, int(helper.Success), true))
	Delete(ctx, pch.service.Delete)
}
