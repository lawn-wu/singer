package model

import (
	"singer/global"
)

type Writer struct {
	global.Model
	Name    string `json:"name" form:"name" gorm:"column:name"`          // 名称
	Profile string `json:"profile" form:"profile" gorm:"column:profile"` // 简介
}

func (Writer) TableName() string {
	return "writer"
}
