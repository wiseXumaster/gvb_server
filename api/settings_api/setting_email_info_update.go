package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
)

func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithSuccess(c)
}
