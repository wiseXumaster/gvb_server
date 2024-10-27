package adverts_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

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
