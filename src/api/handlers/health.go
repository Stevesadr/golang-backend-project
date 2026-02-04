package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(200,"System is ready")
}

func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(200, "working post")
}

func (h *HealthHandler) HealthById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("worked by id %v",id))
}
