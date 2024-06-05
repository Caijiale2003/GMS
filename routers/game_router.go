package routers

import (
	"GameManageSystem/api"
	"GameManageSystem/middleware"
)

func (router RouterGroup) GameRouter() {
	app := api.ApiGroup.GameApi
	router.POST("/game_create", middleware.JwtAdmin(), app.GameCreateView)
	router.GET("/game_search", middleware.JwtSign(), app.GameSearchView)
	router.GET("/game", middleware.JwtSign(), app.GameListView)
	router.POST("/game_enroll", middleware.JwtSign(), app.GameEnrollView)
	router.PUT("/game_revise", middleware.JwtAdmin(), app.GameReviseView)
	router.DELETE("/game_delete", middleware.JwtAdmin(), app.GameDeleteView)
	router.GET("/game_search_student", middleware.JwtAdmin(), app.GameSearchStudentView)
	router.GET("/game_search_self", middleware.JwtSign(), app.GameSearchSelfView)
}
