package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings/:name", settingsApi.SettingsInfoView)
	r.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
	//r.GET("settings", settingsApi.SettingsInfoView)
	//r.PUT("settings", settingsApi.SettingsInfoUpdateView)
	//r.GET("settings_email", settingsApi.SettingsEmailInfoView)
	//r.PUT("settings_email", settingsApi.SettingsEmailInfoUpdateView)
}
