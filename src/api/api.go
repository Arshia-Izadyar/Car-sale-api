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
		property := v1.Group("/property", middleware.Authentication(cfg))
		country := v1.Group("/country", middleware.Authentication(cfg))
		city := v1.Group("/city", middleware.Authentication(cfg))
		year := v1.Group("/year", middleware.Authentication(cfg))
		color := v1.Group("/color", middleware.Authentication(cfg))
		file := v1.Group("/file", middleware.Authentication(cfg))
		gearbox := v1.Group("/gearbox", middleware.Authentication(cfg))
		company := v1.Group("/company", middleware.Authentication(cfg))
		carType := v1.Group("/car-type", middleware.Authentication(cfg))
		carModel := v1.Group("/car-model", middleware.Authentication(cfg))
		carPrice := v1.Group("/car-price", middleware.Authentication(cfg))
		carYear := v1.Group("/car-year", middleware.Authentication(cfg))
		carColor := v1.Group("/car-color", middleware.Authentication(cfg))
		carFile := v1.Group("/car-file", middleware.Authentication(cfg))
		carProperty := v1.Group("/car-property", middleware.Authentication(cfg))
		carComment := v1.Group("/car-comment", middleware.Authentication(cfg))

		router.PersianYearRouter(year, cfg)
		router.CarTypeRouter(carType, cfg)
		router.ColorRouter(color, cfg)
		router.CompanyRouter(company, cfg)
		router.FileRouter(file, cfg)
		router.CityRouter(city, cfg)
		router.PropertyCategoryRouter(propertyCategory, cfg)
		router.CountryRouter(country, cfg)
		router.PropertyRouter(property, cfg)
		router.GearboxRouter(gearbox, cfg)
		router.CarModelRouter(carModel, cfg)
		router.CarModelYearRouter(carYear, cfg)
		router.CarModelPriceRouter(carPrice, cfg)
		router.CarModelColorRouter(carColor, cfg)
		router.CarModelPropertyRouter(carProperty, cfg)
		router.CarModelCommentRouter(carComment, cfg)
		router.CarModelFileRouter(carFile, cfg)

		r.Static("/static/", "../../uploads")
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
