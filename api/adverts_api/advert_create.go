package adverts_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入广告标题" structs:"title"`      // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` // 图片
	//IsShow bool   `json:"is_show" binding:"isbool" msg:"请选择是否展示广告" structs:"is_show"`  // 是否展示
	IsShow bool `json:"is_show" msg:"请选择是否展示广告" structs:"is_show"` // 是否展示
}

// AdvertCreateView 创建广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "广告参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertsApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest

	//绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		//res.FailWithCode(res.ArgumentError, c)
		return
	}

	//重复的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	//数据库执行添加
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error

	//数据库执行添加失败
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}

	res.OkWithMessage("添加广告成功", c)
}
