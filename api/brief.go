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

/*
保存草稿
同时支持修改和创建
如果brief_id 不为空，则更新简报
*/
func SaveBrief() {

}

// 发送给指定人
// 一份简报可以发送给多人？
// 简报发送后置为已发送？
func SendBrief() {

}

// 我的简报
func MyBriefs(c *gin.Context) {

}

// 收到的简报
func ReceiveBriefs(c *gin.Context) {

}

// 导出简报
func ExportBrief() {

}
