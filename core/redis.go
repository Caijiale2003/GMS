package core

import (
	"GameManageSystem/global"
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// ConnectRedis 默认连接0号库
func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	//限制连接所需要的时长
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	//测试连接
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	log.Println("redis连接成功！")
	return rdb
}
