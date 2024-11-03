package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserDetailRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

// UserDetailView 用户信息
// @Tags 用户管理
// @Summary 通过id展示用户信息
// @Description 通过id展示用户信息
// @Router /api/user_detail [get]
// @Param token header string  true  "token"
// @Param data body UserDetailRequest  true  "用户id"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserDetailView(c *gin.Context) {
	var req UserDetailRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if err := c.ShouldBind(&req); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var user models.UserModel
	err := global.DB.Debug().Take(&user, req.ID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	if ctype.Role(claims.Role) != ctype.PermissionAdmin {
		user.UserName = ""
	}

	res.OkWithData(user, c)
}
