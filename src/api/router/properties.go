package router

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/handler"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategoryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)
	r.POST("/create", h.CreatePropertyCategory)
	r.GET("/get/:id", h.GetPropertyCategory)
	r.PUT("/update/:id", h.UpdatePropertyCategory)
	r.DELETE("/delete/:id", h.DeletePropertyCategory)
}

func PropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyHandler(cfg)
	r.POST("/create", h.CreateProperty)
	r.GET("/get/:id", h.GetProperty)
	r.PUT("/update/:id", h.UpdateProperty)
	r.DELETE("/delete/:id", h.DeleteProperty)
}
