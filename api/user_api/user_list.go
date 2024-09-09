package user_api

import (
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/service/common"
	"GVB_server/utils/desens"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {

		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		//脱敏
		users = append(users, user)

	}

	res.OkWithList(users, count, c)

}
