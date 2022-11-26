package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"singer/config"
)

var (
	Logger *zap.Logger
	Config config.Server
	Db     *gorm.DB
	DbList map[string]*gorm.DB
)
