package global

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"` // 主键ID
	CreatedAt time.Time      `json:"-"`                    // 创建时间
	UpdatedAt time.Time      `json:"-"`                    // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
