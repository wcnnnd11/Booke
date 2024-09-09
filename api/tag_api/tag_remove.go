package advert_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)

}
