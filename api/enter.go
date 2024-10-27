package api

import (
	"gvb_server/api/adverts_api"
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertsApi  adverts_api.AdvertsApi
}

var ApiGroupApp = new(ApiGroup)
