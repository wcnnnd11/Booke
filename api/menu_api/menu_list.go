package menu_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MenuResponse]}
// MenuListView 菜单列表
func (MenuApi) MenuListView(c *gin.Context) {
	// 获取查询参数
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 获取搜索关键词
	key := c.Query("key")

	// 查询菜单列表
	var menuList []models.MenuModel
	query := global.DB.Order("sort desc")

	// 如果有搜索关键词，进行模糊查询
	if key != "" {
		query = query.Where("title LIKE ? OR path LIKE ?", "%"+key+"%", "%"+key+"%")
	}

	query.Find(&menuList)

	// 提取菜单ID
	var menuIDList []uint
	for _, menu := range menuList {
		menuIDList = append(menuIDList, menu.ID)
	}

	// 查询关联的 Banner
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id IN ?", menuIDList)

	// 构建返回数据
	var menus = make([]MenuResponse, 0)
	for _, model := range menuList {
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}

	// 返回结果
	res.OkWithList(menus, int64(len(menus)), c)
}
