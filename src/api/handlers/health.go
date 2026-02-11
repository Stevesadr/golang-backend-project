package handlers

import (
	"fmt"
	"net/http"

	"github.com/Stevesadr/golang-backend-project/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(200, helper.GenerateResponse("System is ready",true,0))
}

func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(200,helper.GenerateResponse("working post", true, 0))
}

func (h *HealthHandler) HealthById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, helper.GenerateResponse(fmt.Sprintf("worked by id %v",id), true, 0))
}
