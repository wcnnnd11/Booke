package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) LogV2Router() {
	app := api.ApiGroupApp.LogV2Api
	router.GET("logs/v2", middleware.JwtAdmin(), app.LogListView)
	router.DELETE("logs/v2", middleware.JwtAdmin(), app.LogRemoveView)
	router.GET("logs/v2/read", middleware.JwtAdmin(), app.LogReadView)
}
