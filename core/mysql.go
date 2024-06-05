package core

import (
	"GameManageSystem/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() *gorm.DB {
	// 1、判断host是否配置
	if global.Config.Mysql.Host == "" {
		log.Panic("未配置mysql，取消gorm连接")
		return nil
	}

	// 2、设置mysqllogger的等级
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLogger = logger.Default.LogMode(logger.Info)

	// 3、通过gorm.Open()函数连接到MySQL数据库，并设置日志记录器为上一步配置的mysqlLogger
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	// 4、设置其他
	sqlDB, _ := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最多可容纳
	sqlDB.SetMaxOpenConns(100)
	// 连接最大复用时间，不能超过mysql的wait_timeout
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	// 5、返回db
	log.Println("数据库连接成功！")
	return db
}
