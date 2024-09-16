package redis_ser

import (
	"GVB_server/global"
	"strconv"
)

const lookPrefix = "look"

// Look 浏览某一篇文章
func Look(id string) error {

	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	num++
	err := global.Redis.HSet(lookPrefix, id, num).Err()

	return err
}

// GetLook 获取某一篇文章下的点赞数
func GetLook(id string) int {
	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	return num
}

// GetLookInfo 取出浏览量数据
func GetLookInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(lookPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo

}

func LookClear() {
	global.Redis.Del(lookPrefix)
}
