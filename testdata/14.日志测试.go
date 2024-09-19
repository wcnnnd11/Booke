package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/plugins/log_stash"
	"fmt"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
	log := log_stash.New("192.168.100.158", "xxxx")
	log.Error(fmt.Sprintf("%s 你好啊", "璃江"))
}
