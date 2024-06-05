package game_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/server/common"
	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
	ID string `json:"id" binding:"required" msg:"请输入查询的学号"`
}

func (GameApi) GameSearchStudentView(c *gin.Context) {
	var re SearchRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var team models.EnrollModel
	var cr models.PageInfo
	list, count, _ := common.ComList(models.EnrollModel{}, common.Option{
		PageInfo: cr,
		Where:    global.DB.Take(&team, "id = ?", re.ID),
	})
	res.OkWithList(list, count, c)

}
