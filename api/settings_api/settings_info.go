package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

//var SettingsMap map[string]

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
	//c.JSON(200, gin.H{"msg": "xxx"})
	//res.Ok(map[string]string{"int": "haha"}, "xxx", c)
	//res.FailWithCode(res.SettingsError, c)
	//res.OkWithData(global.Config.SiteInfo, c)
}
