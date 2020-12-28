package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CreateBrief(c *gin.Context) {
	var service service.TaskBriefService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.CreateBrief()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
