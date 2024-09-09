package tag_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var TagList []models.TagModel
	count := global.DB.Find(&TagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 如果这个标签下有文章，怎么办
	global.DB.Delete(&TagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)

}
