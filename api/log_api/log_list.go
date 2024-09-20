package log_api

import (
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/plugins/log_stash"
	"GVB_server/service/common"
	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Leave `form:"level"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	if cr.Sort == "" {
		cr.Sort = "created_at"
	}
	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
