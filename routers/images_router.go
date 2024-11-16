package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.GET("images", app.ImageListView)
	router.GET("images_names", app.ImageNameListView)
	router.POST("images", app.ImageUploadView)
	router.POST("image", middleware.JwtAuth(), app.ImageUploadDataView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)

}
