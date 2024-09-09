package routers

import "GVB_server/api"

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.GET("menus_names", app.MenuNameListView)
	router.GET("menus/:id", app.MenuDetailView)
	router.PUT("menus/:id", app.MenuUpdateView)
	router.DELETE("menus/:id", app.MenuRemoveView)
}
