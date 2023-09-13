package handler

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Company
// @Accept json
// @produces json
// @Param Request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.Response{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/company/create [post]
// @Security AuthBearer
func (pch *CompanyHandler) CreateCompany(ctx *gin.Context) {
	Create[dto.CreateCompanyRequest, dto.CompanyResponse](ctx, pch.service.Create)

}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/get/{id} [get]
// @Security AuthBearer
func (pch *CompanyHandler) GetCompany(ctx *gin.Context) {
	Get[dto.CompanyResponse](ctx, pch.service.GetById)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.Response{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/update/{id} [put]
// @Security AuthBearer
func (pch *CompanyHandler) UpdateCompany(ctx *gin.Context) {
	Update[dto.UpdateCompanyRequest, dto.CompanyResponse](ctx, pch.service.Update)
}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/delete/{id} [delete]
// @Security AuthBearer
func (pch *CompanyHandler) DeleteCompany(ctx *gin.Context) {
	Delete(ctx, pch.service.Delete)
}
