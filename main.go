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

	//router := gin.Default()
	//
	//// 配置CORS中间件，这里允许来自http://localhost:5173的跨域访问，可根据实际需求修改源地址
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"http://localhost:5173"},
	//	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders: []string{"Content-Type", "Authorization"},
	//}))

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
	// 启动 HTTPS 服务器
	//err := router.RunTLS(addr, "/path/to/your_certificate.crt", "/path/to/your_key.key")
	//if err != nil {
	//	global.Log.Fatalf("Failed to start HTTPS server: %s", err.Error())
	//}
}
