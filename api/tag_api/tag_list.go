package advert_api

import (
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OkWithList(list, count, c)
}
