package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("articles/:id", app.ArticleDetailView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)

}
