package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 获取菜单列表
// @Tags 菜单管理
// @Summary 获取菜单列表
// @Description 查询所有菜单并按指定顺序返回，包含菜单的基本信息和关联的图片信息
// @Produce json
// @Success 200 {object} res.Response{data=[]menus_api.MenuResponse}
// @Router /api/menus [get]
func (MenusApi) MenuListView(c *gin.Context) {
	// 使用 Preload 查询菜单及其关联的 Banner
	var menuList []models.MenuModel
	err := global.DB.Preload("Banners").Order("sort desc").Find(&menuList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询菜单列表失败", c)
		return
	}

	// 构建响应数据
	var menus []MenuResponse
	for _, menu := range menuList {
		var banners = make([]Banner, 0)
		for _, banner := range menu.Banners {
			banners = append(banners, Banner{
				ID:   banner.ID,
				Path: banner.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}

	// 返回结果
	res.OkWithData(menus, c)
}
