package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	//res.Ok(map[string]string{"int": "haha"}, "xxx", c)
	res.FailWithCode(res.SettingsError, c)
}
