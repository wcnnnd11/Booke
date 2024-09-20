package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	router.GET("logs", app.LogListView)
	router.DELETE("logs", middleware.JwtAdmin(), app.LogRemoveView)
}
