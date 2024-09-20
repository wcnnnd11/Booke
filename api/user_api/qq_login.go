package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/plugins/qq"
	"GVB_server/utils"
	"GVB_server/utils/jwts"
	"GVB_server/utils/pwd"
	"GVB_server/utils/random"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	//根据openID判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token=?", openID).Error
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		//不存在，就注册
		hashPwd := pwd.HashPwd(random.RandString(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,  //qq登录，邮箱+密码
			Password:   hashPwd, //随机生成16位密码
			Avatar:     qqInfo.Avatar,
			Addr:       addr, //根据IP算地址
			Token:      openID,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}

	}
	//登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		res.FailWithMessage("token生成失败", c)
		return
	}

	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	res.OkWithData(token, c)
}
