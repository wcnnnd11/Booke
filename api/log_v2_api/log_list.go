package log_v2_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	log_stash "GVB_server/plugins/log_stash_v2"
	"GVB_server/service/common"
	"github.com/gin-gonic/gin"
	"time"
)

type LogListRequest struct {
	models.PageInfo
	Level    log_stash.Level   `json:"level" form:"level"`       // 日志查询的等级
	Type     log_stash.LogType `json:"type" form:"type"`         // 日志的类型   1 登录日志  2 操作日志  3 运行日志
	IP       string            `json:"ip" form:"ip"`             // 根据ip查询
	UserID   uint              `json:"userID" form:"userID"`     // 根据用户id查询
	Addr     string            `json:"addr" form:"addr"`         // 感觉地址查询
	Date     string            `json:"date" form:"date"`         // 查某一天的，格式是年月日
	Status   *bool             `json:"status" form:"status"`     // 登录状态查询  true  成功  false 失败
	UserName string            `json:"userName" form:"userName"` // 查用户名
}

// LogListView 日志列表
// @Tags 日志管理V2
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogListRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/v2 [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log_stash.LogModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("状态错误", c)
		return
	}
	var query = global.DB.Where("")
	if cr.Date != "" {
		_, dateTimeErr := time.Parse("2006-01-02", cr.Date)
		if dateTimeErr != nil {
			res.FailWithMessage("时间格式错误", c)
			return
		}
		query.Where("date(created_at) = ?", cr.Date)
	}
	if cr.Status != nil {
		query.Where("status = ?", cr.Status)
	}

	_list, count, _ := common.ComList(log_stash.LogModel{
		Type:     cr.Type,
		Level:    cr.Level,
		IP:       cr.IP,
		Addr:     cr.Addr,
		UserID:   cr.UserID,
		UserName: cr.UserName,
	}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    query,
		Likes:    []string{"title", "user_name"},
	})
	res.OkWithList(_list, count, c)
}
