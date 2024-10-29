package api

import (
	"gvb_server/api/adverts_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menus_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertsApi  adverts_api.AdvertsApi
	MenusApi    menus_api.MenusApi
}

var ApiGroupApp = new(ApiGroup)
