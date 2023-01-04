package validate

import (
	internal "singer/app/validate/internal"
	"singer/global/request"
)

type SaveTagValidate request.SaveTagRequest

func (*SaveTagValidate) GetError(err string) string {
	msg := internal.TransferMessage(TagErrorMessage, []string{"TagName"})
	return internal.GetError(msg, err)
}

var TagErrorMessage = map[string]map[string]string{
	"TagName": {"required": "请输入歌曲类型标签", "max": "歌曲类型标签输入最大长度为10"},
}
