package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// UserInfoView 用户详细信息接口
// @Tags 用户管理
// @Summary 获取用户详细信息
// @Description 获取用户详细信息，包括积分、签名等
// @Router /api/user_info [get]
// @Param token header string true "token"
// @Produce json
// @Success 200 {object} res.Response{data=models.UserModel}
func (UserApi) UserInfoView(c *gin.Context) {
	// 获取当前用户的 claims
	_claims, exists := c.Get("claims")
	if !exists {
		res.FailWithMessage("用户未登录", c)
		return
	}
	claims := _claims.(*jwts.CustomClaims)

	// 根据 token 中的用户 ID 查询用户信息
	var user models.UserModel
	err := global.DB.Debug().Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 如果不是管理员，隐藏部分敏感信息
	if ctype.Role(claims.Role) != ctype.PermissionAdmin {
		user.UserName = "" // 隐藏用户名
	}
	// 返回完整的用户信息
	res.OkWithData(user, c)
}
