package menus_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" msg:"请完善菜单名称" structs:"menu_title"`
	MenuTitleEn   string      `json:"menu_title_en" msg:"请完善菜单英文名称" structs:"menu_title_en"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

// MenuCreateView 创建菜单
// @Tags 菜单管理
// @Summary 创建一个新菜单
// @Description 创建菜单时可以选择关联图片和设置菜单的各项信息
// @Accept json
// @Produce json
// @Param menu_body body menus_api.MenuRequest true "菜单信息"
// @Success 200 {object} res.Response{data=interface{},msg=string} "菜单添加成功"
// @Router /api/menus [post]
func (MenusApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 创建菜单模型实例
	menuModel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	// 启动事务来确保菜单创建和图片关联的一致性
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建菜单记录
		if err := tx.Create(&menuModel).Error; err != nil {
			return err
		}

		// 如果没有图片需要关联，直接返回
		if len(cr.ImageSortList) == 0 {
			return nil // 直接结束事务，菜单已创建，无需关联图片
		}

		// 构建菜单与图片的关联记录
		var menuBannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			menuBannerList = append(menuBannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}

		// 将关联数据插入数据库
		if err := tx.Create(&menuBannerList).Error; err != nil {
			return err // 如果插入失败，事务回滚
		}

		return nil
	})

	// 检查事务结果
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败或图片关联失败", c)
		return
	}

	res.OkWithMessage("菜单添加成功", c)
}
