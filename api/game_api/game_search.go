package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameSearchRequest struct {
	Name string `json:"name" binding:"required" msg:"请输入比赛名称"`
}

// GameSearchView 查询比赛
func (GameApi) GameSearchView(c *gin.Context) {
	// 参数绑定校验
	var re GameSearchRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 校验数据库中是否已经创建该比赛
	var game models.GameModel
	err = global.DB.Take(&game, "name = ?", re.Name).Error
	if err != nil {
		res.FailWithMessage("未找到该比赛！请您仔细核对比赛名称是否正确", c)
		return
	}

	// 找到该比赛，然后返回信息
	c.JSON(http.StatusOK, gin.H{
		"比赛名称":   game.Name,
		"比赛主办方":  game.Organizer,
		"比赛开始时间": game.StartTime,
		"比赛结束时间": game.EndTime,
		"比赛地点":   game.Address,
		"比赛奖品":   game.Prize,
	})
	res.OkWithData("创建成功", c)

}
