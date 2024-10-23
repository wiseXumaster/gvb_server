package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	//var cr config.SiteInfo
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	//err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		//var info config.SiteInfo
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.Jwt = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithSuccess(c)

	//global.Config.SiteInfo = cr
	//err = core.SetYaml()
	//if err != nil {
	//	global.Log.Error(err)
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//res.OkWithSuccess(c)
}
