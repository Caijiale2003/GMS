package routers

import (
	"GameManageSystem/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	//设置gin模式
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//创建路由组
	apiRouterGroup := router.Group("/api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.UserRouter()
	routerGroupApp.GameRouter()
	routerGroupApp.ScoreRouter()
	return router
}
