package service

import (
	"errors"
	"gorm.io/gorm"
	"singer/app/model"
	"singer/global"
	"singer/global/request"
)

type SongService struct {
}

var SongServiceApp = new(SongService)

// Enter
// @Description: 录入歌曲信息
// @receiver *SongService
// @param params
// @return err
func (*SongService) Enter(params request.SongEnterRequest) (err error) {
	tx := global.Db.Begin()

	// 录入作词人信息
	var writer model.Writer
	if errors.Is(tx.Where("name = ?", params.Writer).First(&writer).Error, gorm.ErrRecordNotFound) {
		writer = model.Writer{
			Name: params.Writer,
		}
		if err = tx.Save(&writer).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	// 录入作曲人信息
	var composer model.Composer
	if errors.Is(tx.Where("name = ?", params.Composer).First(&composer).Error, gorm.ErrRecordNotFound) {
		composer = model.Composer{
			Name: params.Composer,
		}
		if err = tx.Save(&composer).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	// 录入歌手信息
	var singer model.Singer
	if errors.Is(tx.Where("name = ?", params.Singer).First(&singer).Error, gorm.ErrRecordNotFound) {
		singer = model.Singer{
			Name: params.Singer,
		}
		if err = tx.Save(&singer).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	// 录入歌曲信息
	var song model.Song
	if errors.Is(tx.Where("name = ? AND singer_id = ?", params.SongName, singer.ID).First(&song).Error, gorm.ErrRecordNotFound) {
		song = model.Song{
			Name:       params.SongName,
			WriterId:   writer.ID,
			ComposerId: composer.ID,
			SingerId:   singer.ID,
			Profile:    params.Profile,
		}
		if err = tx.Save(&song).Error; err != nil {
			tx.Rollback()
			return
		}

		// 录入歌词
		var lyric model.Lyric
		if errors.Is(tx.Where("song_id = ?", song.ID).First(&lyric).Error, gorm.ErrRecordNotFound) {
			lyric = model.Lyric{
				SongId:  song.ID,
				Content: params.Lyric,
			}
			if err = tx.Save(&lyric).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	}

	tx.Commit()
	return
}
