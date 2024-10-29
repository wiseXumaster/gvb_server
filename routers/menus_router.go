package routers

import (
	"gvb_server/api"
)

func (r RouterGroup) MenusRouter() {
	MenusApi := api.ApiGroupApp.MenusApi
	r.POST("menus", MenusApi.MenuCreateView)
}
