package adverts_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// AdvertsUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param id path string true "广告ID"
// @Param data body AdvertRequest true "广告参数"
// @Router /api/adverts/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertsApi) AdvertsUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	//// 使用 Updates 方法更新多个字段
	//err = global.DB.Model(&advert).Updates(models.AdvertModel{
	//	Title:  cr.Title,
	//	Href:   cr.Href,
	//	Images: cr.Images,
	//	IsShow: cr.IsShow,
	//}).Error
	//直接转换为map
	crmaps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(crmaps).Error

	if err != nil {
		res.FailWithMessage("更新广告失败", c)
		return
	}

	res.OkWithMessage("广告更新成功", c)

}
