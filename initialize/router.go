package initialize

import (
	"singer/global"
	"singer/router"

	"github.com/gin-gonic/gin"
)

// Routers
// @Description: 初始化总路由
// @return *gin.Engine
func Routers() *gin.Engine {
	Router := gin.Default()

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		router.InitPullRouter(PublicGroup)
	}

	global.Logger.Info("router register success")
	return Router
}
