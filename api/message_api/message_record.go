package message_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户ID"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)

	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range _messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}

	// 点开消息，里面的每一条消息，都从未读变成已读

	res.OkWithData(messageList, c)
}
