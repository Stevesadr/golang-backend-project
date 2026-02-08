package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("token")
		if apiKey != "1"{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "token is invalid",
			})
			return 
		}
		ctx.Next()  
	}
}