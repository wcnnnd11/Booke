package api

import (
	"GVB_server/api/advert_api"
	"GVB_server/api/images_api"
	"GVB_server/api/menu_api"
	"GVB_server/api/settings_api"
	"GVB_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserAPi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
