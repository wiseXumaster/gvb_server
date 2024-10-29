package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageAirListView 获取图片简要列表
// @Tags 图片管理
// @Summary 获取图片简要列表
// @Description 获取分页的图片简要列表，仅返回图片的基本信息（id, path, name）
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageResponse}
// @Router /api/images_air [get]
func (ImagesApi) ImageAirListView(c *gin.Context) {
	var imageList []ImageResponse
	err := global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询图片轻量列表失败", c)
		return
	}
	res.OkWithData(imageList, c)
}
