package user_api

import "github.com/gin-gonic/gin"

func (UserApi) LoginInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":       "请输入学号",
		"name":     "请输入姓名",
		"password": "请输入密码",
		"gender":   "请输入性别",
		"academy":  "请输入学院",
		"major":    "请输入专业",
	})
}
