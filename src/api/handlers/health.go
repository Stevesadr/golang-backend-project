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

// @Summary System ready
// @Description This handler is just hear to tell system is ready
// @Accept json
// @Produce json
// @Tags Health
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(200, helper.GenerateResponse("System is ready",true,0))
}

// @Summary Post System ready
// @Description This handler is just hear to tell post system is ready
// @Accept json
// @Produce json
// @Tags Health
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/health/ [post]
func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(200,helper.GenerateResponse("working post", true, 0))
}

// @Summary Get Param
// @Description This handler want to get id from param 
// @Accept json
// @Produce json
// @Tags Health
// @Param id path string true "User id" 
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/health/{id} [get]
func (h *HealthHandler) HealthById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, helper.GenerateResponse(fmt.Sprintf("worked by id %s",id), true, 0))
}
