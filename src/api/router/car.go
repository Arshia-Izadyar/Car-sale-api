package router

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/handler"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/gin-gonic/gin"
)

func GearboxRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewGearboxHandler(cfg)
	r.POST("/create", h.CreateGearbox)
	r.GET("/get/:id", h.GetGearbox) // /gearbox
	r.PUT("/update/:id", h.UpdateGearbox)
	r.DELETE("/delete/:id", h.DeleteGearbox)
}

func CarTypeRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarTypeHandler(cfg)
	r.POST("/create", h.CreateCarType)
	r.GET("/get/:id", h.GetCarType)
	r.PUT("/update/:id", h.UpdateCarType) // /car-type
	r.DELETE("/delete/:id", h.DeleteCarType)
}

func CompanyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCompanyHandler(cfg)
	r.POST("/create", h.CreateCompany)
	r.GET("/get/:id", h.GetCompany) // /company
	r.PUT("/update/:id", h.UpdateCompany)
	r.DELETE("/delete/:id", h.DeleteCompany)
}

func CarModelRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelHandler(cfg)
	r.POST("/create", h.CreateCarModel)
	r.GET("/get/:id", h.GetCarModel)
	r.PUT("/update/:id", h.UpdateCarModel)
	r.DELETE("/delete/:id", h.DeleteCarModel)
}

func CarModelPriceRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelPriceHandler(cfg)
	r.POST("/create", h.CreateCarModelPrice)
	r.GET("/get/:id", h.GetCarModelPrice)
	r.PUT("/update/:id", h.UpdateCarModelPrice)
	r.DELETE("/delete/:id", h.DeleteCarModelPrice)
}

func CarModelYearRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelYearHandler(cfg)
	r.POST("/create", h.CreateCarModelYear)
	r.GET("/get/:id", h.GetCarModelYear)
	r.PUT("/update/:id", h.UpdateCarModelYear)
	r.DELETE("/delete/:id", h.DeleteCarModelYear)
}

func CarModelColorRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelColorHandler(cfg)
	r.POST("/create", h.CreateCarModelColor)
	r.GET("/get/:id", h.GetCarModelColor)
	r.PUT("/update/:id", h.UpdateCarModelColor)
	r.DELETE("/delete/:id", h.DeleteCarModelColor)
}

func CarModelFileRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelFileHandler(cfg)
	r.POST("/create", h.CreateCarModelFile)
	r.GET("/get/:id", h.GetCarModelFile)
	r.PUT("/update/:id", h.UpdateCarModelFile)
	r.DELETE("/delete/:id", h.DeleteCarModelFile)
}

func CarModelPropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelPropertyHandler(cfg)
	r.POST("/create", h.CreateCarModelProperty)
	r.GET("/get/:id", h.GetCarModelProperty)
	r.PUT("/update/:id", h.UpdateCarModelProperty)
	r.DELETE("/delete/:id", h.DeleteCarModelProperty)
}

func CarModelCommentRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCarModelCommentHandler(cfg)
	r.POST("/create", h.CreateCarModelComment)
	r.GET("/get/:id", h.GetCarModelComment)
	r.PUT("/update/:id", h.UpdateCarModelComment)
	r.DELETE("/delete/:id", h.DeleteCarModelComment)
}
