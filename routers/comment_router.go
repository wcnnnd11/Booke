package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments", app.CommentListView)
	router.GET("comments/:id", app.CommentDiggView)
	router.DELETE("comments/:id", app.CommentRemoveView)

	router.GET("comments/articles", app.CommentByArticleListView)

}
