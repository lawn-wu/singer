package validate

import (
	internal "singer/app/validate/internal"
	"singer/global/request"
)

type ComposerEnterValidate request.ComposerEnterRequest

func (*ComposerEnterValidate) GetError(err string) string {
	msg := internal.TransferMessage(ComposerErrorMessage, []string{"Name", "Profile"})
	return internal.GetError(msg, err)
}

var ComposerErrorMessage = map[string]map[string]string{
	"Name":    {"required": "请输入作曲人", "max": "作曲人输入最大长度为10"},
	"Profile": {"required": "请输入歌曲简介", "max": "歌曲简介输入最大长度为255"},
}
