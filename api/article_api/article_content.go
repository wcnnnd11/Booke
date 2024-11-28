package article_api

import (
	"GVB_server/models/res"
	"GVB_server/service/es_ser"
	"github.com/gin-gonic/gin"
)

type ArticleContentResponse struct {
	// 只返回文件的内容
	Content string `json:"content" form:"content"`
}

type ArticleContentRequest struct {
	ID string `json:"id" form:"id"`
}

func (ArticleApi) ArticleContentByIdView(c *gin.Context) {
	var cr ArticleContentRequest
	// 绑定查询参数
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 使用新方法根据 ID 获取文章内容
	content, err := es_ser.CommDetailByID(cr.ID)
	if err != nil {
		if err.Error() == "文章不存在" {
			res.FailWithCode(res.Error, c)
			return
		}
		res.FailWithMessage(err.Error(), c)
		return
	}

	// 返回文章内容
	res.OkWithData(ArticleContentResponse{
		Content: content,
	}, c)
}
