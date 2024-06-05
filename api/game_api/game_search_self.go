package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/server/common"
	"GameManageSystem/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (GameApi) GameSearchSelfView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var team models.EnrollModel
	var cr models.PageInfo
	list, count, _ := common.ComList(models.EnrollModel{}, common.Option{
		PageInfo: cr,
		Where:    global.DB.Take(&team, "id = ?", claims.UserID),
	})
	res.OkWithList(list, count, c)
}
