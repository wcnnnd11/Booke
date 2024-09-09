package images_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`                // 图片路径
	Name string `gorm:"size:38" json:"name"` // 图片名称
}

// ImageNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {

	var imageList []ImageResponse

	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)

	res.OkWithData(imageList, c)

}
