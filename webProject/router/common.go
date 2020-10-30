package router

import (
	"github.com/gin-gonic/gin"
	"webProject/api"
)

func InitCommonRouter(group *gin.Engine) {
	decisionRouter := group.Group("common")
	{
		decisionRouter.GET("showPlat", api.ShowPlat)
	}
}