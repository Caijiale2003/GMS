package routers

import "GameManageSystem/api"

func (router RouterGroup) UserRouter() {
	app := api.ApiGroup.UserApi
	router.GET("/sign_view", app.SignInfoView)
	router.GET("/login_view", app.LoginInfoView)
	router.POST("/login", app.UserLoginView)
	router.POST("/sign", app.UserSignView)
}
