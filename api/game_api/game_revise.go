package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"github.com/gin-gonic/gin"
)

func (GameApi) GameReviseView(c *gin.Context) {
	var re GameCreateRequest
	//传参
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var game models.GameModel
	err = global.DB.Debug().Take(&game, "name = ?", re.Name).Error
	if err != nil {
		res.FailWithMessage("该比赛不存在", c)
		return
	}

	err = global.DB.Where("name = ?", re.Name).Updates(models.GameModel{
		Organizer: re.Organizer,
		StartTime: re.StartTime,
		EndTime:   re.EndTime,
		Address:   re.Address,
		Prize:     re.Prize}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("比赛修改失败", c)
		return
	}
	res.OkWithMessage("比赛修改成功", c)
	return

}
