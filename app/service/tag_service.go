package service

import (
	"errors"
	"gorm.io/gorm"
	"singer/app/model"
	"singer/global"
	"singer/global/request"
)

type TagService struct {
}

var TagServiceApp = new(TagService)

// SaveTag
// @Description: 添加歌曲类型标签
// @receiver *TagService
// @param params
// @return err
func (*TagService) SaveTag(params request.SaveTagRequest) (err error) {
	var tag model.Tag
	if errors.Is(global.Db.Where("name = ?", params.TagName).First(&tag).Error, gorm.ErrRecordNotFound) {
		tag = model.Tag{
			Name: params.TagName,
		}
		if err = global.Db.Save(&tag).Error; err != nil {
			return
		}
	}
	return
}
