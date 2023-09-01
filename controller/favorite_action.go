package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/service"
)

func FavoriteAction(c *gin.Context) {

	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	var action string
	if actionType == "1" {
		action = "点赞"
	} else if actionType == "2" {
		action = "取消点赞"
	}

	err := service.FavoriteAction(videoId, actionType)
	if err != nil {
		log.Println("controller: " + err.Error())
		FavoriteActionErrorResponse(c, action+"失败")
		c.Abort()
		return
	}

	FavoriteActionSuccessResponse(c, action+"成功")
}

func FavoriteActionErrorResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: msg})
}

func FavoriteActionSuccessResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: msg})
}
