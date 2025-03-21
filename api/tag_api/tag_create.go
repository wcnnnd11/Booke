package tag_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}

	res.OkWithMessage("添加标签成功", c)
}
