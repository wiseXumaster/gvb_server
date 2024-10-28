package adverts_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest true "广告ID列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertsApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var advertList []models.AdvertModel
	var count int64
	global.DB.Find(&advertList, cr.IDList).Count(&count)

	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}

	global.DB.Delete(advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 条广告", count), c)

}
