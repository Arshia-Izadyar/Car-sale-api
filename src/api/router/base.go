package router

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/handler"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/gin-gonic/gin"
)

func CountryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)
	r.POST("/create", h.CreateCountry)
	r.GET("/get/:id", h.GetCountry)
	r.PUT("/update/:id", h.UpdateCountry)
	r.DELETE("/delete/:id", h.DeleteCountry)
}

func CityRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCityHandler(cfg)
	r.POST("/create", h.CreateCity)
	r.GET("/get/:id", h.GetCity)
	r.PUT("/update/:id", h.UpdateCity)
	r.DELETE("/delete/:id", h.DeleteCity)
}

func ColorRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewColorHandler(cfg)
	r.POST("/create", h.CreateColor)
	r.GET("/get/:id", h.GetColor)
	r.PUT("/update/:id", h.UpdateColor)
	r.DELETE("/delete/:id", h.DeleteColor)
}

func PersianYearRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPersianYearHandler(cfg)
	r.POST("/create", h.CreatePersianYear)
	r.GET("/get/:id", h.GetPersianYear)
	r.PUT("/update/:id", h.UpdatePersianYear)
	r.DELETE("/delete/:id", h.DeletePersianYear)
}
