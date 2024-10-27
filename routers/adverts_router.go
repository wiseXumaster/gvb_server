package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) AdvertsRouter() {
	advertsApi := api.ApiGroupApp.AdvertsApi
	r.POST("adverts", advertsApi.AdvertCreateView)
	r.GET("adverts", advertsApi.AdvertsListView)
	r.PUT("adverts/:id", advertsApi.AdvertsUpdateView)
	r.DELETE("adverts", advertsApi.AdvertRemoveView)
}
