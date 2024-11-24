package message_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/service/common"
	"github.com/gin-gonic/gin"
)

type MessageUserRecordRequest struct {
	models.PageInfo
	SendUserID uint `json:"sendUserID" form:"sendUserID" binding:"required"`
	RevUserID  uint `json:"revUserID" form:"revUserID" binding:"required"`
}

// MessageUserRecordView 两个用户之间的聊天记录
// @Tags 消息管理
// @Summary 两个用户之间的聊天记录
// @Description 两个用户之间的聊天记录
// @Router /api/message_users/record [get]
// @Param token header string  true  "token"
// @Param data query MessageUserRecordRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageUserRecordView(c *gin.Context) {
	var cr MessageUserRecordRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}

	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", cr.SendUserID, cr.RevUserID, cr.SendUserID, cr.RevUserID),
	})

	res.OkWithList(list, count, c)
}
