package router

import (
	"github.com/gin-gonic/gin"
	"singer/app/api"
)

func InitPullRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("song")
	songApi := api.SongApiExample
	{
		baseRouter.POST("enter", songApi.Enter)
	}
	return baseRouter
}
