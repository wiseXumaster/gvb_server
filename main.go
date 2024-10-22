package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	//global.Log.Println("ok")
	//global.Log.Warnln("warn")
	//global.Log.Error("warn")
	//global.Log.Infof("warn")
	//logrus.Warn("hh")
	// 连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
	//fmt.Println(global.Config)
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	router.Run(addr)
}
