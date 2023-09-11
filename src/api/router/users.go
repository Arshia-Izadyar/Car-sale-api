package router

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/handler"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)
	r.POST("/get-otp", h.SendOtp)
	r.POST("/login/phone", h.RegisterLoginByPhone)
	r.POST("/register/phone", h.RegisterLoginByPhone)
	r.POST("/register/username", h.RegisterByUsername)
	r.POST("/login/username", h.LoginByUsername)
}
