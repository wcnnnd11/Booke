package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	// 连接es
	global.ESClient = core.EsConnect()

	/*
		[{"tag":"python","article_count":2,"article_list":[]}]
	*/

	agg := elastic.NewTermsAggregation().Field("tags")

	//agg.SubAggregation("articles_id", elastic.NewTermsAggregation().Field("_id"))  //// 查_id测试
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))

	query := elastic.NewBoolQuery()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
	//fmt.Println(string(result.Aggregations["tags"]))   // 查_id测试
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}
	fmt.Println(tagList)

}
