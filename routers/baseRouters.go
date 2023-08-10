package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/controller"
	"tiktok/util"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")

	baseRouters.GET("/feed", controller.Feed)

	baseRouters.GET("/test", func(c *gin.Context) {
		// 查询数据库
		videoList := []util.Author{}

		util.DB.Find(&videoList)

		c.JSON(http.StatusOK, gin.H{
			"video": videoList,
		})
	})
}
