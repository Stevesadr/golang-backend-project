package middlewares

import (
	"net/http"

	"github.com/Stevesadr/golang-backend-project/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc{
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateResponseWithError(gin.H{
				"error": err.Error()  }, false, 1, err))
			return 
		} 
		ctx.Next()
	}
}