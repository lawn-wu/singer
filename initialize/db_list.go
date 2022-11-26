package initialize

import (
	"gorm.io/gorm"
	"singer/config"
	"singer/global"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.Config.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "sqlserver":
			dbMap[info.AliasName] = GormSqlServerByConfig(config.SqlServer{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.Db = sysDB
	}
	global.DbList = dbMap
}
