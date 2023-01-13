package service

import (
	"errors"
	"gorm.io/gorm"
	"singer/app/model"
	"singer/global"
	"singer/global/request"
)

type ComposerService struct {
}

var ComposerServiceApp = new(ComposerService)

// Enter
// @Description: 录入作曲人信息
// @receiver *ComposerService
// @param params
// @return err
func (*ComposerService) Enter(params request.ComposerEnterRequest) (err error) {
	tx := global.Db.Begin()

	// 录入作曲人信息
	var composer model.Composer
	if errors.Is(tx.Where("name = ?", params.Name).First(&composer).Error, gorm.ErrRecordNotFound) {
		composer = model.Composer{
			Name: params.Name,
			Profile: params.Profile,
		}
		if err = tx.Save(&composer).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}
