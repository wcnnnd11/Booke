package log_stash

import (
	"GVB_server/global"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip, token string) *Log {
	// 解析token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   "内网",
		userId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLeave, content)
}

func (l Log) Info(content string) {
	l.send(InfoLeave, content)

}

func (l Log) Warn(content string) {
	l.send(WarnLeave, content)

}

func (l Log) Error(content string) {
	l.send(ErrorLeave, content)

}

func (l Log) send(level Leave, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
	//fmt.Println(l.ip, l.userId, l.addr, content, level)
}

//func Debug(ip string, content string) {
//	std.Debug(ip, content)
//}
//func Info(ip string, content string) {
//	std.Debug(ip, content)
//}
//func Warn(ip string, content string) {
//	std.Debug(ip, content)
//}
//func Error(ip string, content string) {
//	std.Debug(ip, content)
//}
