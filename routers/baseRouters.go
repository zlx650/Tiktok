package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/controller"
	"tiktok/middleware"
)

func BaseRoutersInit(r *gin.Engine) {

	r.Static("static", "./static")
	r.GET("/ping/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"ping": "success",
		})
	})

	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed/", controller.Feed)

		baseRouters.POST("/publish/action/", middleware.JWTMiddleWare(), controller.PublishVideoController)
		baseRouters.POST("/user/register/", controller.Register)
		baseRouters.POST("/user/login/", controller.Login)
		baseRouters.GET("/user/", middleware.JWTMiddleWare(), controller.UserInfo)
		baseRouters.GET("/publish/list/", middleware.JWTMiddleWare(), controller.GetPublishList)

		//baseRouters.GET("/user", controller.UserInfo)
	}

}
