package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.GET("messages_users", app.MessageUserListView)
	router.GET("messages_users/record", app.MessageUserRecordView)
	router.GET("messages_users/user", app.MessageUserListByUserView)

	router.GET("messages", middleware.JwtAuth(), app.MessageListView)
	router.POST("messages", middleware.JwtAuth(), app.MessageCreateView)
	router.GET("messages_all", app.MessageListAllView)
	router.GET("messages_record", middleware.JwtAuth(), app.MessageRecordView)

}
