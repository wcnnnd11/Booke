package main

import (
	"GVB_server/global"
	"context"
	"github.com/sirupsen/logrus"
)

func (DemoModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExist 查询索引是否存在
func (demo DemoModel) IndexExist() bool {
	exists, err := client.IndexExists(demo.Index()).Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (demo DemoModel) CreateIndex() error {
	if demo.IndexExist() {
		// 有索引
		demo.RemoveIndex()
	}
	// 无索引
	// 就创建索引
	createIndex, err := client.
		CreateIndex(demo.Index()).
		BodyString(demo.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建索引")
		return err
	}
	logrus.Infof("索引 %s 创建成功", demo.Index())
	return nil
}

// RemoveIndex 删除索引
func (demo DemoModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}
