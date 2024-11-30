package routers

import (
	"GVB_server/api"
)

func (router RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	router.GET("chat_groups", app.ChatGroupView)
	router.GET("chat_groups_records", app.ChatListView)
	router.DELETE("chat_groups_remove", app.ChatRemoveView)

}
