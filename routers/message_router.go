package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.GET("messages_users/me", middleware.JwtAuth(), app.MessageUserListByMeView)
	router.GET("messages_users/record/me", middleware.JwtAuth(), app.MessageUserRecordByMeView)

	router.GET("messages_users", middleware.JwtAuth(), app.MessageUserListView)
	router.GET("messages_users/record", middleware.JwtAuth(), app.MessageUserRecordView)
	router.GET("messages_users/user", middleware.JwtAuth(), app.MessageUserListByUserView)

	router.GET("messages", middleware.JwtAuth(), app.MessageListView)
	router.POST("messages", middleware.JwtAuth(), app.MessageCreateView)
	router.GET("messages_all", app.MessageListAllView)
	router.GET("messages_record", middleware.JwtAuth(), app.MessageRecordView)

}
