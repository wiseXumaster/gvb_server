package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) AdvertsRouter() {
	advertsApi := api.ApiGroupApp.AdvertsApi
	r.POST("adverts", advertsApi.AdvertCreateView)
}
