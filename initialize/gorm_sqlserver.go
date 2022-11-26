package initialize

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"singer/config"
	"singer/global"
	"singer/initialize/internal"
)

// GormSqlServer 初始化SQL Server数据库
func GormSqlServer() *gorm.DB {
	s := global.Config.SqlServer
	if s.Dbname == "" {
		return nil
	}
	sqlServerConfig := sqlserver.Config{
		DSN:               s.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(sqlServerConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

// GormSqlServerByConfig 初始化SQL Server数据库用过传入配置
func GormSqlServerByConfig(s config.SqlServer) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}
	sqlServerConfig := sqlserver.Config{
		DSN:               s.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(sqlServerConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
