package user_api

import (
	"GVB_server/global"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name,select(c)" binding:"required" msg:"请输入昵称"` // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`          // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`            // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`                // 权限  1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return

}
