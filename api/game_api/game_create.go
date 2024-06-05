package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/utils/jwts"
	"github.com/gin-gonic/gin"
)

type GameCreateRequest struct {
	Name      string `json:"name" binding:"required" msg:"请输入比赛名称"`
	Organizer string `json:"organizer" binding:"required" msg:"请输入比赛主办方"`
	StartTime string `json:"start_time" binding:"required" msg:"请输入比赛开始时间"`
	EndTime   string `json:"end_time" binding:"required" msg:"请输入比赛结束时间"`
	Address   string `json:"address" binding:"required" msg:"请输入比赛地点"`
	Prize     string `json:"prize" binding:"required" msg:"请输入比赛奖品"`
}

// GameCreateView 创建比赛
func (GameApi) GameCreateView(c *gin.Context) {
	// 参数绑定校验
	var re GameCreateRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 校验数据库中是否已经创建该比赛
	var game models.GameModel
	err = global.DB.Take(&game, "name = ?", re.Name).Error
	if err == nil {
		res.FailWithMessage("该比赛已经创建！如您需要更改比赛信息，请在左侧选择修改比赛信息", c)
		return
	}
	// 没有该比赛，然后创建
	game = models.GameModel{
		Name:      re.Name,
		Organizer: re.Organizer,
		StartTime: re.StartTime,
		EndTime:   re.EndTime,
		Address:   re.Address,
		Prize:     re.Prize,
		Creator:   claims.Username,
	}
	err = global.DB.Create(&game).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("创建失败", c)
		return
	}
	res.OkWithData("创建成功", c)
}
