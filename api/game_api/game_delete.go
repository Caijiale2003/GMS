package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	GameName string `json:"game_name" binding:"required" msg:"请输入比赛名称"`
}

func (GameApi) GameDeleteView(c *gin.Context) {
	// 参数绑定校验
	var re DeleteRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 删除单条记录
	global.DB.Where("name = ?", re.GameName).Delete(models.GameModel{})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}

	res.OkWithMessage("删除成功", c)
}
