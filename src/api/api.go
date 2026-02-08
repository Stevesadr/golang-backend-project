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

func InitServer(){
	cfg := config.GetConfig()

	r := gin.Default() 

	valid, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		valid.RegisterValidation("mobile",validations.IranianMobileNumberValidation, true)
		valid.RegisterValidation("password", validations.PasswordValidation, true)
	}

	// r.Use(middlewares.TestingMiddleware())
	r.Use(middlewares.LimitByRequest())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("test")
		routers.HealthRouter(health)
		routers.Testing(test)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}