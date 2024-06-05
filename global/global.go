package global

import (
	"GameManageSystem/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	Log         *logrus.Logger
	DB          *gorm.DB
	MysqlLogger logger.Interface
	Redis       *redis.Client
)
