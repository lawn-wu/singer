package model

import (
	"singer/global"
)

type Tag struct {
	global.Model
	Name string `json:"name" form:"name" gorm:"column:name"` // 名称
}

func (Tag) TableName() string {
	return "tag"
}
