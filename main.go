package main

import (
	"GVB_server/core"
	_ "GVB_server/docs"
	"GVB_server/flag"
	"GVB_server/global"
	"GVB_server/routers"
)

// @title gvb_API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
	//连接redis
	global.Redis = core.ConnectRedis()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Info("程序GVB_server运行在:", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
