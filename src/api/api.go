package api

import (
	"github.com/Stevesadr/golang-backend-project/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	r := gin.Default() 

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)
	}

	r.Run(":5005")
}