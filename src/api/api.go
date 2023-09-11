package api

import (
	"fmt"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/router"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/validators"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init(cfg *config.Config) {
	logger := logging.NewLogger(cfg)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	registerRoutes(r, cfg)
	registerValidators()

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatal(err, logging.General, logging.Api, "cant run server", nil)
	}
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		users := v1.Group("/users")
		router.UserRouter(users, cfg)
	}
}

func registerValidators() {
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("password", validators.PassWordValidator, true)
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
	}
}
