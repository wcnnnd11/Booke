package settings_api

import (
	"GVB_server/global"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		emailInfo := global.Config.Email
		emailInfo.Password = "******"
		res.OkWithData(emailInfo, c)
	case "qq":
		qq := global.Config.QQ
		qq.Key = "******"
		res.OkWithData(qq, c)
	case "qiniu":
		qiniu := global.Config.QiNiu
		qiniu.SecretKey = "******"
		res.OkWithData(qiniu, c)
	case "jwt":
		jwt := global.Config.Jwt
		jwt.Secret = "******"
		res.OkWithData(jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
