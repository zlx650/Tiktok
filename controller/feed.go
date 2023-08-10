package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/util"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []util.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	// 查询数据库获取video信息

	videoList := []util.Video{}
	util.DB.Preload("Author").Find(&videoList)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(), // TODO fix feed NextTime
	})
}
