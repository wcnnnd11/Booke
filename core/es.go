package core

import (
	"GVB_server/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.ES.URL()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败%s", err.Error())
	}
	return c
}
