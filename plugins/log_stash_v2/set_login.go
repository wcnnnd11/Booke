package log_stash_v2

import (
	"GVB_server/global"
	"github.com/gin-gonic/gin"
)

// NewSuccessLogin 登录成功的日志
func NewSuccessLogin(c *gin.Context, userID uint, userName string) {
	saveLoginLog("登录成功", "--", userID, userName, true, c)
}

// NewFailLogin 登录失败的日志
func NewFailLogin(title string, userName string, pwd string, c *gin.Context) {
	saveLoginLog(title, pwd, 0, userName, false, c)
}

func saveLoginLog(title string, content string, userID uint, userName string, status bool, c *gin.Context) {
	ip := c.ClientIP()
	addr := getAddr(ip)
	global.DB.Create(&LogModel{
		IP:       ip,
		Addr:     addr,
		Title:    title,
		Content:  content,
		UserID:   userID,
		UserName: userName,
		Status:   status,
		Type:     LoginType,
	})
}
