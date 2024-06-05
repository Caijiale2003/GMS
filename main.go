package main

import (
	"GameManageSystem/core"
	"GameManageSystem/global"
	"GameManageSystem/routers"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitMysql()
	// 连接redis
	global.Redis = core.ConnectRedis()
	//初始化路由
	router := routers.InitRouter()
	//返回ip地址
	addr := global.Config.System.Addr()
	//运行
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
