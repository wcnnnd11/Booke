package user_api

import (
	"GVB_server/models"
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/service/common"
	"GVB_server/utils/desens"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 获取请求中的 `role` 参数
	role := c.DefaultQuery("role", "")
	var roleInt int
	if role != "" {
		// 将 `role` 参数转换为整型
		roleInt, _ = strconv.Atoi(role)
	}

	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Likes:    []string{"nick_name", "user_name"},
		Role:     roleInt, // 传递角色过滤条件
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	time.Sleep(1 * time.Second)

	res.OkWithList(users, count, c)
}
