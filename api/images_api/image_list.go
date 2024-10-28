package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/models"
	"gvb_server/service/common"
)

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 获取图片列表
// @Description 获取分页的图片列表
// @Param page query int false "当前页码"
// @Param page_size query int false "每页大小"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
		//ModelType: "BannerModel",
	})
	// 确保 page 和 limit 的值合理

	//res.OkWithData(gin.H{"count": count, "list": imageList}, c)
	res.OkWithList(list, count, c)
	return
}
