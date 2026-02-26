package api

import (
	"fmt"

	"github.com/Stevesadr/golang-backend-project/api/middlewares"
	"github.com/Stevesadr/golang-backend-project/api/routers"
	"github.com/Stevesadr/golang-backend-project/api/validations"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config){

	r := gin.Default() 

	RegisterValidators() 

	// r.Use(middlewares.TestingMiddleware())
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))

	RegisterRoute(r)

	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterValidators(){

	valid, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		valid.RegisterValidation("mobile",validations.IranianMobileNumberValidation, true)
		valid.RegisterValidation("password", validations.PasswordValidation, true)
	}

}

func RegisterRoute(r *gin.Engine){
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("test")
		routers.HealthRouter(health)
		routers.Testing(test)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config){
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "Backend of project"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}