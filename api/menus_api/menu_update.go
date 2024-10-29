package menus_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// MenuUpdateView 更新菜单
// @Tags 菜单管理
// @Summary 更新菜单
// @Description 根据菜单 ID 更新菜单信息，包括关联的 Banner 图片列表
// @Param id path string true "菜单 ID"
// @Param menu_body body MenuRequest true "更新菜单信息的请求体"
// @Produce json
// @Success 200 {object} res.Response{msg=string} "菜单更新成功"
// @Failure 400 {object} res.Response{msg=string} "请求体解析错误"
// @Failure 404 {object} res.Response{msg=string} "菜单不存在"
// @Failure 500 {object} res.Response{msg=string} "菜单更新失败"
// @Router /api/menus/{id} [put]
func (MenusApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	id := c.Param("id")

	// 启动事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 查找菜单
		var menuModel models.MenuModel
		if err := tx.Take(&menuModel, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				res.FailWithMessage("菜单不存在", c)
			} else {
				global.Log.Error(err)
				res.FailWithMessage("查询菜单失败", c)
			}
			return err
		}

		// 清空旧的 Banner 关联
		if err := tx.Model(&menuModel).Association("Banners").Clear(); err != nil {
			global.Log.Error(err)
			res.FailWithMessage("清空旧的 Banner 关联失败", c)
			return err
		}

		// 添加新的 Banner 关联
		if len(cr.ImageSortList) > 0 {
			var bannerList []models.MenuBannerModel
			for _, sort := range cr.ImageSortList {
				bannerList = append(bannerList, models.MenuBannerModel{
					MenuID:   menuModel.ID,
					BannerID: sort.ImageID,
					Sort:     sort.Sort,
				})
			}

			if err := tx.Create(&bannerList).Error; err != nil {
				global.Log.Error(err)
				res.FailWithMessage("创建关联失败", c)
				return err
			}
		}

		// 更新普通字段
		maps := structs.Map(&cr)
		filteredMaps := make(map[string]interface{})
		for key, value := range maps {
			if key != "ImageSortList" && tx.Migrator().HasColumn(&menuModel, key) {
				filteredMaps[key] = value
			}
		}

		if err := tx.Model(&menuModel).Updates(filteredMaps).Error; err != nil {
			global.Log.Error(err)
			res.FailWithMessage("修改菜单失败", c)
			return err
		}

		return nil
	})

	// 检查事务结果
	if err != nil {
		return // 错误处理已在事务内部完成，无需进一步处理
	}

	res.OkWithMessage("修改菜单成功", c)
}
