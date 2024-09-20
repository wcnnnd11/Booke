package data_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
	"time"
)

type DataCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateResponse struct {
	DateList  []string `json:"date_list"`
	LoginDate []int    `json:"login_date"`
	SignDate  []int    `json:"sign_date"`
}

var dateCount = map[string]int{}

func (DataApi) SevenLoginView(c *gin.Context) {
	var loginDataCount, signDateCount []DataCount

	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDataCount)

	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)
	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int

	now := time.Now()
	for _, i2 := range loginDataCount {
		loginDateCountMap[i2.Date] = i2.Count
	}
	for _, i2 := range loginDataCount {
		signDateCountMap[i2.Date] = i2.Count
	}
	var dateList []string
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginDateCountMap[day]
		signCount := signDateCountMap[day]
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}

	res.OkWithData(DateResponse{
		DateList:  dateList,
		LoginDate: loginCountList,
		SignDate:  signCountList,
	}, c)
}
