package redis_ser

import (
	"GVB_server/global"
	"strconv"
)

const diggPrefix = "digg"

// Digg 点赞某一篇文章
func Digg(id string) error {

	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	num++
	err := global.Redis.HSet(diggPrefix, id, num).Err()

	return err
}

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	return num
}

// GetDiggInfo 每隔一段时间同步点赞数据到es
func GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo

}

func DiggClear() {
	global.Redis.Del(diggPrefix)
}
