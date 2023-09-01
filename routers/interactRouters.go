package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok/authmiddleware"
	"tiktok/controller"
)

func InteractRoutersInit(r *gin.Engine) {

	interactRouters := r.Group("douyin")
	{
		interactRouters.POST("/favorite/action", authmiddleware.JWTMiddleWare(), controller.FavoriteAction)
	}

}
