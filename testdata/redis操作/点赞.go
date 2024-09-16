package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/service/redis_ser"
	"fmt"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()

	global.Redis = core.ConnectRedis()

	redis_ser.Digg("VXEA9JEBKWTZixGuebWs")
	fmt.Println(redis_ser.GetDigg("VXEA9JEBKWTZixGuebWs"))

	fmt.Println(redis_ser.GetDiggInfo())
	//redis_ser.DiggClear()
}
