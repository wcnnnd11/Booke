package advert_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)

}
