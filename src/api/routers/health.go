package routers

import (
	"github.com/Stevesadr/golang-backend-project/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup){
	h := handlers.NewHealthHandler()
	
	r.GET("/", h.Health )
	r.GET("/:id", h.HealthById )
	r.POST("/", h.HealthPost)
}