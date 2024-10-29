package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// MenuAirList 获取菜单简要列表
// @Tags 菜单管理
// @Summary 获取菜单简要列表
// @Description 获取仅包含菜单 ID、标题和路径的简要列表
// @Produce json
// @Success 200 {object} res.Response{data=[]MenuNameResponse}
// @Router /api/menus_air [get]
func (MenusApi) MenuAirList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "menu_title as title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
