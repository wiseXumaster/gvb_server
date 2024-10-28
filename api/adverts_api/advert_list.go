package adverts_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/models"
	"gvb_server/service/common"
	"strings"
)

// AdvertsListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param page query int false "当前页码"
// @Param page_size query int false "每页大小"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertsApi) AdvertsListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//默认只显示true
	referer := c.GetHeader("Referer")
	isShow := true

	//如果是管理员全显示
	if strings.Contains(referer, "admin") {
		isShow = false
	}

	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)

}
