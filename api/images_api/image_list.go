package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/models"
	"gvb_server/service/common"
)

//图片列表
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
	})
	// 确保 page 和 limit 的值合理

	//res.OkWithData(gin.H{"count": count, "list": imageList}, c)
	res.OkWithList(list, count, c)
	return
}
