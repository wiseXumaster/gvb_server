package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// MenuDetailView 获取菜单详情
// @Tags 菜单管理
// @Summary 获取菜单详情
// @Description 根据菜单 ID 获取菜单的详细信息，包括关联的 Banner 列表
// @Accept json
// @Produce json
// @Param id path string true "菜单 ID"
// @Success 200 {object} res.Response{data=MenuResponse} "成功返回菜单详情数据"
// @Failure 404 {object} res.Response{msg=string} "菜单不存在"
// @Router /api/menus/{id} [get]
func (MenusApi) MenuDetailView(c *gin.Context) {
	// 先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerModel.ID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
	return

}
