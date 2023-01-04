package api

import (
	"github.com/gin-gonic/gin"
	"singer/app/service"
	"singer/app/validate"
	"singer/global/request"
	"singer/global/response"
)

type TagApi struct {
}

// SaveTag
// @Description: 添加歌曲类型标签
// @receiver *TagApi
// @param c
func (*TagApi) SaveTag(c *gin.Context) {
	// 验证请求参数
	var params validate.SaveTagValidate
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage(params.GetError(err.Error()), c)
		return
	}

	if err := service.TagServiceApp.SaveTag(request.SaveTagRequest(params)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
