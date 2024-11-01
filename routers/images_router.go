package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("images", imagesApi.ImageUploadView)
	r.GET("images", imagesApi.ImageListView)
	r.DELETE("images", imagesApi.ImageRemoveView)
	r.PUT("images", imagesApi.ImageUpdateView)
	r.GET("images_air", imagesApi.ImageAirListView)
	//r.GET("settings", settingsApi.SettingsInfoView)
	//r.PUT("settings", settingsApi.SettingsInfoUpdateView)
	//r.GET("settings_email", settingsApi.SettingsEmailInfoView)
	//r.PUT("settings_email", settingsApi.SettingsEmailInfoUpdateView)
}
