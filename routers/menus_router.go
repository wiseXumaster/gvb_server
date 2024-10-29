package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) MenusRouter() {
	MenusApi := api.ApiGroupApp.MenusApi
	r.POST("menus", MenusApi.MenuCreateView)
	r.GET("menus", MenusApi.MenuListView)
	r.GET("menus_air", MenusApi.MenuAirList)
	r.PUT("menus/:id", MenusApi.MenuUpdateView)
	r.DELETE("menus", MenusApi.MenuRemoveView)
	r.GET("menus/:id", MenusApi.MenuDetailView)
}
