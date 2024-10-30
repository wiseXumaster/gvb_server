package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
	//Router *gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	//router.Use(cors.Default())
	// 路由分组
	apiRouterGroup := router.Group("api")
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	routerGroupApp := RouterGroup{apiRouterGroup}
	//SettingsRouter(router)
	//routerGroupApp := RouterGroup{router}
	// 路由分层
	// 系统配置api
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertsRouter()
	routerGroupApp.MenusRouter()

	//b站弹幕:传入router,再给group会不会好些
	//router.Run()
	return router
}
