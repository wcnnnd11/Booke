package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/plugins/log_stash"
	"GVB_server/plugins/log_stash_v2"
	"GVB_server/utils"
	"GVB_server/utils/jwts"
	"GVB_server/utils/pwd"
	"fmt"
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

	log := log_stash.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name=? or email=?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 用户名不存在
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		log_stash_v2.NewFailLogin("用户名不存在", cr.UserName, cr.Password, c) // 调用登录失败日志方法
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s %s", cr.UserName, cr.Password))
		log_stash_v2.NewFailLogin("用户名密码错误", cr.UserName, cr.Password, c) // 调用登录失败日志方法
		res.FailWithMessage("用户名或密码错误", c)
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
		log.Error(fmt.Sprintf("token 生成失败 %s", err.Error()))
		res.FailWithMessage("token生成失败", c)
		return
	}

	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(c.ClientIP(), token)
	log.Info("登录成功")

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	log_stash_v2.NewSuccessLogin(c) // 调用登录成功日志方法

	res.OkWithData(token, c)
}
