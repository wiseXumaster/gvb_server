package menus_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// MenuRemoveView 删除菜单
// @Tags 菜单管理
// @Summary 删除菜单
// @Description 根据提供的菜单 ID 列表，删除对应的菜单项，并清空关联的 Banners
// @Accept json
// @Produce json
// @Param menu_body body models.RemoveRequest true "删除菜单的请求体，包含菜单 ID 列表"
// @Success 200 {object} res.Response{msg=string} "返回删除成功的消息，包括删除的菜单数量"
// @Failure 400 {object} res.Response{msg=string} "请求参数错误"
// @Failure 500 {object} res.Response{msg=string} "删除菜单或清空关联失败"
// @Router /api/menus [delete]
func (MenusApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	var count int64
	// 先查询菜单是否存在
	global.DB.Find(&menuList, cr.IDList).Count(&count)
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 启动事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 清空每个菜单的 Banners 关联
		for _, menu := range menuList {
			if err := tx.Model(&menu).Association("Banners").Clear(); err != nil {
				global.Log.Error(err)
				return err
			}
		}

		// 删除菜单
		if err := tx.Delete(&menuList).Error; err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})

	// 检查事务结果
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
