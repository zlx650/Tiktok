package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed", controller.Feed)
		baseRouters.POST("/publish/action/", controller.PublishVideoController)
	}

}
