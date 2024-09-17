package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/service/es_ser"
)

func main() {
	core.InitConf()
	core.InitLogger()
	global.ESClient = core.EsConnect()
	es_ser.DeleteFullTextByArticleID("VHH785EBKWTZixGuzLUF")

}
