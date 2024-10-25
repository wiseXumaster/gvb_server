package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel
	var count int64
	global.DB.Find(&imageList, cr.IDList).Count(&count)

	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}

	global.DB.Delete(imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)

}
