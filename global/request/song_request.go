package request

type SongEnterRequest struct {
	SongName string `json:"songName" form:"songName" binding:"required,max=20"` // 歌曲名称
	Writer   string `json:"writer" form:"writer" binding:"required,max=10"`     // 作词人
	Composer string `json:"composer" form:"composer" binding:"required,max=10"` // 作曲人
	Singer   string `json:"singer" form:"singer" binding:"required,max=10"`     // 演唱人
	Profile  string `json:"profile" form:"profile" binding:"max=255"`           // 歌曲简介
	Lyric    string `json:"lyric" form:"lyric" binding:"required"`              // 歌曲歌词
}
