package api

import (
	"GVB_server/api/advert_api"
	"GVB_server/api/article_api"
	"GVB_server/api/chat_api"
	"GVB_server/api/comment_api"
	"GVB_server/api/data_api"
	"GVB_server/api/digg_api"
	"GVB_server/api/images_api"
	"GVB_server/api/log_api"
	"GVB_server/api/log_v2_api"
	"GVB_server/api/menu_api"
	"GVB_server/api/message_api"
	"GVB_server/api/role_api"
	"GVB_server/api/settings_api"
	"GVB_server/api/tag_api"
	"GVB_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserAPi     user_api.UserApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	LogV2Api    log_v2_api.LogApi
	DataApi     data_api.DataApi
	RoleApi     role_api.RoleApi
}

var ApiGroupApp = new(ApiGroup)
