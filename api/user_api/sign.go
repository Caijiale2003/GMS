package user_api

import "github.com/gin-gonic/gin"

func (UserApi) SignInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":       "请输入学号",
		"password": "请输入密码",
	})
}
