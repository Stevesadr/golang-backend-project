package api

import (
	"fmt"

	"github.com/Stevesadr/golang-backend-project/api/routers"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	cfg := config.GetConfig()

	r := gin.Default() 

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}