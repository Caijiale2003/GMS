package score_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"github.com/gin-gonic/gin"
)

type ScoreEntryRequest struct {
	GameName string `json:"game_name" binding:"required" msg:"请输入比赛名称"`
	ID       string `json:"id" binding:"required" msg:"请输入学号"`
	Score    int    `json:"score" binding:"required" msg:"请输入得分"`
}

func (ScoreApi) ScoreEntryView(c *gin.Context) {
	// 参数绑定校验
	var re ScoreEntryRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 查看是否有此比赛
	var game models.GameModel
	err = global.DB.Take(&game, "name = ?", re.GameName).Error
	if err != nil {
		res.FailWithMessage("该比赛不存在，请检查比赛名称是否正确", c)
		return
	}

	// 校验数据库中是否已经有了该条记录
	var score models.ScoreModel
	err = global.DB.Take(&score, "game_name = ? && id = ?", re.GameName, re.ID).Error
	if err == nil {
		res.FailWithMessage("该记录已存在，请移步修改", c)
		return
	}
	// 有该比赛，然后创建
	err = global.DB.Create(&models.ScoreModel{
		GameName: re.GameName,
		ID:       re.ID,
		Score:    re.Score,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("写入失败", c)
		return
	}
	res.OkWithData("写入成功", c)
}
