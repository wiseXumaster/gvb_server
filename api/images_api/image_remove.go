package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

// ImageRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 根据提供的图片 ID 列表批量删除图片
// @Param image_body body models.RemoveRequest true "要删除的图片 ID 列表"
// @Produce json
// @Success 200 {object} res.Response{msg=string} "图片删除成功，返回删除数量"
// @Failure 400 {object} res.Response{msg=string} "请求体解析错误"
// @Failure 404 {object} res.Response{msg=string} "文件不存在"
// @Router /api/images [delete]
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
