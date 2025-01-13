package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/plugins/log_stash_v2"
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
		// 可以选择在此记录参数绑定失败的日志，使用登录日志或其他方式
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? OR email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		res.FailWithMessage("用户名或密码错误", c)
		// 记录登录失败日志
		log_stash_v2.NewFailLogin("用户名不存在", cr.UserName, cr.Password, c)
		return
	}

	// 校验密码
	isValid := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isValid {
		res.FailWithMessage("用户名或密码错误", c)
		// 记录登录失败日志
		log_stash_v2.NewFailLogin("密码错误", cr.UserName, cr.Password, c)
		return
	}

	// 登录成功，生成 token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token 生成失败", c)
		// 可以选择在此记录 token 生成失败的日志
		return
	}

	// 记录登录成功日志
	log_stash_v2.NewSuccessLogin(c, userModel.ID, userModel.UserName)

	res.OkWithData(token, c)
}
