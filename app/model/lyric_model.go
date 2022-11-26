package model

import (
	"singer/global"
)

type Lyric struct {
	global.Model
	SongId  uint   `json:"songId" form:"songId" gorm:"column:song_id"`   // 歌曲id
	Content string `json:"content" form:"content" gorm:"column:content"` // 内容
}

func (Lyric) TableName() string {
	return "lyric"
}
