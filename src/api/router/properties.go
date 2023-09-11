package router

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/handler"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategoryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)
	r.POST("/create", h.Create)
}
