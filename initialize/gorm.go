package initialize

import (
	"gorm.io/gorm"
	"singer/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "sqlserver":
		return GormSqlServer()
	default:
		return GormMysql()
	}
}
