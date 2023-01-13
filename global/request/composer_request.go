package request

type ComposerEnterRequest struct {
	Name    string `json:"name" form:"name" binding:"required,max=10"` // 作曲人
	Profile string `json:"profile" form:"profile" binding:"max=255"`   // 作曲人简介
}
