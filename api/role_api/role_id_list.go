package role_api

import (
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

type OptionResponse struct {
	Label string `json:"label"`
	Valve int    `json:"value"`
}

func (RoleApi) RoleIDListView(c *gin.Context) {
	res.OkWithData([]OptionResponse{
		{"管理员", 1},
		{"平台用户", 2},
		{"游客", 3},
	}, c)
}
