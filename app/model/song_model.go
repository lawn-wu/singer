package model

import (
	"singer/global"
	"strings"
)

type Song struct {
	global.Model
	Name       string `json:"name" form:"name" gorm:"column:name"`                    // 名称
	WriterId   uint   `json:"writerId" form:"writerId" gorm:"column:writer_id"`       // 作词人id
	ComposerId uint   `json:"composerId" form:"composerId" gorm:"column:composer_id"` // 作曲人id
	SingerId   uint   `json:"singerId" form:"singerId" gorm:"column:singer_id"`       // 演唱人id
	Profile    string `json:"profile" form:"profile" gorm:"column:profile"`           // 简介
	TagIds     string `json:"tagIds" form:"tagIds" gorm:"column:tag_ids"`             // 标签id组合

	// 关联歌词模型
	Lyric Lyric
}

func (*Song) TableName() string {
	return "song"
}

// GetSplitLyric
// @Description: 获取分解后的歌词
// @receiver song
// @return lyric
func (song *Song) GetSplitLyric() (lyric []string) {
	if song.Lyric.ID <= 0 {
		global.Db.Preload("Lyric").First(song)
	}

	lyric = strings.Split(song.Lyric.Content, "\r\n")
	return
}