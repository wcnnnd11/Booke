package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/utils/jwts"
	"GVB_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name=? or email=?", cr.UserName, cr.UserName).Error
	if err != nil {
		//没找到
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		res.FailWithMessage("tkoen生成失败", c)
		return
	}
	res.OkWithData(token, c)

}
