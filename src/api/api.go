package api

import (
	"fmt"

	"github.com/Stevesadr/golang-backend-project/api/middlewares"
	"github.com/Stevesadr/golang-backend-project/api/routers"
	"github.com/Stevesadr/golang-backend-project/api/validations"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config){

	r := gin.Default() 

	RegisterValidators() 

	// r.Use(middlewares.TestingMiddleware())
	r.Use(middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))

	RegisterRoute(r)

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