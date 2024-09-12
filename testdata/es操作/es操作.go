package main

import (
	"GVB_server/core"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"

	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败%s", err.Error())
	}
	return c
}

func init() {
	core.InitConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

// Create 增加索引功能
func Create(data *DemoModel) (err error) {
	IndexResponse, err := client.Index().Index(data.Index()).BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	data.ID = IndexResponse.Id
	return nil
}

// FindList 查看
func FindList(key string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

func FindSourceList(key string, page, limit int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Source(`{"_source": ["title"]}`).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []DemoModel{}
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	fmt.Println(demoList, count)
}

// Update 更新
func Update(id string, data *DemoModel) error {
	_, err := client.
		Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Info("更新demo成功")
	return nil
}

// Remove 批量删除
func Remove(idList []string) (count int, err error) {
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

func main() {
	//DemoModel{}.CreateIndex()
	//Create(&DemoModel{Title: "Golang框架学习", UserID: 2, CreatedAt: time.Now().Format("2006-01-02 15:04:05")})
	list, count := FindList("", 1, 10)
	fmt.Println(list, count)
	//FindSourceList("python", 1, 10) // 搜索似乎失效了
	//Update("2NFx5ZEBwvxzjfHfYr50", &DemoModel{Title: "python学习来"})
	//count, err := Remove([]string{"2dEg5pEBwvxzjfHf5L7t", "2tEh5pEBwvxzjfHfYr5J"})
	//fmt.Println(count, err)
}
