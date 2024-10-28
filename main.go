package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gvb_server/cmd"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/utils"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host localhost:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	option := cmd.Parse()
	if cmd.IsWebStop(option) {
		cmd.SwitchOption(option)
		return
	}

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isbool", utils.IsBool)
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
