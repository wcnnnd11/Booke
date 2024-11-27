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
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 根据 ID 获取文章详情
	model, err := es_ser.ComDetail(cr.ID)
	if err != nil {
		if err.Error() == "文章不存在" {
			res.FailWithCode(res.Error, c) // 返回自定义 NotFound 错误码
			return
		}
		res.FailWithMessage(err.Error(), c)
		return
	}

	// 返回文章内容
	response := ArticleContentResponse{
		Content: model.Content,
	}
	res.OkWithData(response, c)
}
