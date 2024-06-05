package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/utils/jwts"
	"github.com/gin-gonic/gin"
)

type EnrollRequest struct {
	GameName string `json:"game_name" binding:"required" msg:"请输入比赛名称"`
	Teacher  string `json:"teacher" binding:"required" msg:"请输入指导老师"`
	TeamName string `json:"team_name"`
}

func (GameApi) GameEnrollView(c *gin.Context) {
	// 参数绑定校验
	var re EnrollRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 获取登录者信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 校验数据库中是否已经创建该比赛
	var game models.GameModel
	err = global.DB.Take(&game, "name = ?", re.GameName).Error
	if err != nil {
		res.FailWithMessage("该比赛不存在，请检查比赛名称是否正确", c)
		return
	}

	// 校验该学生是否已经报名比赛
	var enroll models.EnrollModel
	err = global.DB.Take(&enroll, "game_name = ? && id = ?", re.GameName, claims.UserID).Error
	if err == nil {
		res.FailWithMessage("您已报名该比赛，请勿重复报名", c)
		return
	}

	// 未报名，报名写表
	err = global.DB.Create(&models.EnrollModel{
		GameName: re.GameName,
		ID:       claims.UserID,
		Name:     claims.Username,
		Teacher:  re.Teacher,
		TeamName: re.TeamName,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("报名失败", c)
		return
	}

	if re.TeamName != "" {
		var team models.TeamModel
		err = global.DB.Take(&team, "team_name = ? && id = ?", re.TeamName, claims.UserID).Error
		if err == nil {
			res.FailWithMessage("队伍已经创建", c)
			return
		}
		err = global.DB.Create(&models.TeamModel{
			TeamName: re.TeamName,
			ID:       claims.UserID,
			Name:     claims.Username,
		}).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建队伍失败", c)
			return
		}
	}

	res.OkWithData("报名成功", c)
}
