package menu_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	// 通过 path 查询菜单
	//id := c.Param("id") // 这个要去了，查连接表直接用 menuModel.ID
	path := c.DefaultQuery("path", "") //从 查询参数 中获取参数
	//path := c.Param("path") //从 URL 路径 中获取动态参数
	var menuModel models.MenuModel
	err := global.DB.Where("path = ?", path).Take(&menuModel).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", menuModel.ID)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}

	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
	return
}
