package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"singer/core"
	"singer/global"
	"singer/initialize"
)

func main() {
	// 初始化viper
	_ = core.Viper()

	// 初始化zap日志库
	global.Logger = core.Zap()
	zap.ReplaceGlobals(global.Logger)

	// gorm连接数据库
	global.Db = initialize.Gorm()
	// 多数据库服务
	initialize.DBList()
	if global.Db != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.Db.DB()
		defer db.Close()
	}

	if !global.Config.System.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	core.RunWindowsServer()
}
