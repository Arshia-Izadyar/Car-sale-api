package api

import (
	"fmt"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/middleware"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/router"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/validators"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/docs"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(cfg *config.Config) {
	logger := logging.NewLogger(cfg)
	r := gin.New()
	r.Use(middleware.StructuredLog(logger), gin.Recovery())
	r.Use(middleware.Limiter())
	r.Use(middleware.Cors(cfg))
	registerRoutes(r, cfg)
	registerValidators()
	registerSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatal(err, logging.General, logging.Api, "cant run server", nil)
	}
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// TODO: add auth
		users := v1.Group("/users")
		router.UserRouter(users, cfg)

		propertyCategory := v1.Group("/property-category", middleware.Authentication(cfg))
		router.PropertyCategoryRouter(propertyCategory, cfg)

		property := v1.Group("/property", middleware.Authentication(cfg))
		router.PropertyRouter(property, cfg)
	}
}

func registerValidators() {
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("password", validators.PassWordValidator, true)
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
	}
}

func registerSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "car sale api"
	docs.SwaggerInfo.Description = "golang api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Server.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
