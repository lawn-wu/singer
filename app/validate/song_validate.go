package validate

import (
	internal "singer/app/validate/internal"
	"singer/global/request"
)

type SongEnterValidate request.SongEnterRequest

func (*SongEnterValidate) GetError(err string) string {
	msg := internal.TransferMessage(BasicErrorMessage, []string{"SongName", "Writer", "Composer", "Singer", "Profile", "Lyric"})
	return internal.GetError(msg, err)
}

var BasicErrorMessage = map[string]map[string]string{
	"SongName": {"required": "请输入歌曲名称", "max": "歌曲名称输入最大长度为20"},
	"Writer":   {"required": "请输入作词人", "max": "作词人输入最大长度为10"},
	"Composer": {"required": "请输入作曲人", "max": "作曲人输入最大长度为10"},
	"Singer":   {"required": "请输入演唱者", "max": "演唱者输入最大长度为10"},
	"Profile":  {"required": "请输入歌曲简介", "max": "歌曲简介输入最大长度为255"},
	"Lyric":    {"required": "请输入歌曲歌词"},
}
