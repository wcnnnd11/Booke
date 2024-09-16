package redis_ser

import (
	"GVB_server/global"
	"GVB_server/utils"
	"fmt"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	fmt.Println(keys)
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
