package log_v2_api

import (
	"GVB_server/global"
	"GVB_server/models/res"
	log_stash "GVB_server/plugins/log_stash_v2"
	"github.com/gin-gonic/gin"
)

type IDRequest struct {
	ID    uint   `json:"id" binding:"required" msg:"请输入日志id"`
	Token string `gorm:"size:256" json:"token"`
}

// LogReadView 日志读取
// @Tags 日志管理V2
// @Summary 日志读取
// @Description 日志读取
// @Description 1. 前端判断这个日志的读取状态，未读就去请求这个接口，让这个日志变成已读的
// @Description 2. 如果是已读状态，就不需要调这个接口了
// @Param data query models.IDRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/v2/read [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogApi) LogReadView(c *gin.Context) {
	var cr IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	var log log_stash.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMessage("日志不存在", c)
		return
	}
	if log.ReadStatus {
		res.OkWithMessage("日志读取成功", c)
		return
	}
	global.DB.Model(&log).Update("readStatus", true)
	res.OkWithMessage("日志读取成功", c)
	return
}
