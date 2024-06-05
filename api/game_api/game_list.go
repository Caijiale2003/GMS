package game_api

import (
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/server/common"
	"github.com/gin-gonic/gin"
)

type GameListRequest struct {
	Name      string `json:"name"`
	Organizer string `json:"organizer"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (GameApi) GameListView(c *gin.Context) {
	var cr models.PageInfo
	list, count, _ := common.ComList(models.GameModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
