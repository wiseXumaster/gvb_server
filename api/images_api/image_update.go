package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required"  msg:"请选择文件"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 更新图片名称
// @Tags 图片管理
// @Summary 更新图片名称
// @Description 根据图片 ID 更新图片的名称
// @Param image_body body images_api.ImageUpdateRequest true "图片 ID 和新的名称"
// @Produce json
// @Success 200 {object} res.Response{msg=string} "图片名称修改成功"
// @Failure 400 {object} res.Response{msg=string} "请求体解析错误或参数无效"
// @Failure 404 {object} res.Response{msg=string} "文件不存在"
// @Router /api/images [put]
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
	return
}
