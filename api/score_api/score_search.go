package score_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/server/common"
	"github.com/gin-gonic/gin"
)

type ScoreSearchRequest struct {
	GameName string `json:"game_name" binding:"required" msg:"请输入比赛名称"`
}

func (ScoreApi) ScoreSearchView(c *gin.Context) {
	// 参数绑定校验
	var re ScoreSearchRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 校验数据库中是否已经创建该比赛
	var score models.ScoreModel
	err = global.DB.Take(&score, "game_name = ?", re.GameName).Error
	if err != nil {
		res.FailWithMessage("未找到该比赛！请您仔细核对比赛名称是否正确", c)
		return
	}
	var temp models.ScoreModel
	var cr models.PageInfo
	list, count, _ := common.ComList(models.ScoreModel{}, common.Option{
		PageInfo: cr,
		Where:    global.DB.Take(&temp, "game_name = ?", re.GameName),
	})
	res.OkWithList(list, count, c)
}
