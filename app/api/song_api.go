package api

import (
	"github.com/gin-gonic/gin"
	"singer/app/service"
	"singer/app/validate"
	"singer/global/request"
	"singer/global/response"
)

type SongApi struct {
}

var SongApiExample = new(SongApi)

// Enter
// @Description: 录入歌曲
// @receiver *SongApi
// @param c
func (*SongApi) Enter(c *gin.Context) {
	// 验证请求参数
	var params validate.SongEnterValidate
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage(params.GetError(err.Error()), c)
		return
	}

	if err := service.SongServiceApp.Enter(request.SongEnterRequest(params)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
