package request

type SaveTagRequest struct {
	TagName string `json:"tagName" form:"tagName" binding:"required,max=10"` // 歌曲类型标签名称
}
