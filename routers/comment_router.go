package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments/:id", app.CommentListView) // 文章下的评论列表
	router.GET("comments", app.CommentListView)     // 评论列表
	router.GET("comments/digg/:id", app.CommentDiggView)
	router.DELETE("comments/:id", app.CommentRemoveView)

	router.GET("comments/articles", app.CommentByArticleListView)

}
