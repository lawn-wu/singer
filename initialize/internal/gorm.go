package internal

import (
	"log"
	"os"
	"singer/global"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbBase interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * 1000 * 1000,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DbBase
	switch global.Config.System.DbType {
	case "mysql":
		logMode = &global.Config.Mysql
	case "sqlserver":
		logMode = &global.Config.SqlServer
	default:
		logMode = &global.Config.Mysql
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
