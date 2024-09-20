package utils

import (
	"GVB_server/global"
)

func PrintSystem() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, i := range ipList {
			global.Log.Infof("程序GVB_server运行在:https://%s:%d/api", i, port)
			global.Log.Infof("gvb_server api文档运行在：http://%s:%d/swagger/index.html#", i, port)
		}
	} else {
		global.Log.Infof("程序GVB_server运行在:https://%s:%d/api", ip, port)
		global.Log.Infof("gvb_server api文档运行在：http://%s:%d/swagger/index.html#", ip, port)

	}
}
