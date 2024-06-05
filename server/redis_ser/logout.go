package redis_ser

import (
	"GameManageSystem/global"
	"time"
)

const prefix = "logout_"

func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if InList(prefix+token, keys) {
		return true
	}
	return false
}
