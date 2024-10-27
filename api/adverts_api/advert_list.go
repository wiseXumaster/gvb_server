package adverts_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/models"
	"gvb_server/service/common"
	"strings"
)

func (AdvertsApi) AdvertsListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	referer := c.GetHeader("Referer")
	isShow := true

	if strings.Contains(referer, "admin") {
		isShow = false
	}

	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)

}
