package api

import (
	"github.com/gin-gonic/gin"
	"singer/app/service"
	"singer/app/validate"
	"singer/global/request"
	"singer/global/response"
)

type ComposerApi struct {
}

var ComposerApiExample = new(ComposerApi)

// Enter
// @Description: 录入作曲人
// @receiver *ComposerApi
// @param c
func (*ComposerApi) Enter(c *gin.Context) {
	// 验证请求参数
	var params validate.ComposerEnterValidate
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage(params.GetError(err.Error()), c)
		return
	}

	if err := service.ComposerServiceApp.Enter(request.ComposerEnterRequest(params)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
