package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middleware"
)

func InteractRoutersInit(r *gin.Engine) {

	interactRouters := r.Group("douyin")
	{
		interactRouters.POST("/favorite/action/", middleware.JWTMiddleWare(), controller.FavoriteAction)
	}

}
