package routers

import (
	"GameManageSystem/api"
	"GameManageSystem/middleware"
)

func (router RouterGroup) ScoreRouter() {
	app := api.ApiGroup.ScoreApi
	router.POST("/score_entry", middleware.JwtAdmin(), app.ScoreEntryView)
	router.GET("/score_search", middleware.JwtSign(), app.ScoreSearchView)
}
