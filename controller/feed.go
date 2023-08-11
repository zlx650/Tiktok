package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/util"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []util.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	// 返回参数
	latestTime := c.DefaultQuery("latest_time", strconv.FormatInt(time.Now().Unix(), 10))
	// token := c.Query("token") // TODO: fix token parameter

	// 参数转换
	postTime, err := util.ConvertTimestampStrToUnix(latestTime)
	if err != nil {
		return
	}

	// 从数据库中取videoList数据
	videoList := []util.Video{}
	util.DB.Preload("Author").Where("post_time < ?", postTime).Order("post_time desc").Limit(30).Find(&videoList)

	// 选出videoList中最早的post_time
	var nextTime int64 = time.Now().Unix()
	if len(videoList) > 0 {
		for _, video := range videoList {
			videoTime := video.PostTime.Unix()
			if videoTime < nextTime {
				nextTime = video.PostTime.Unix()
			}
		}
	}

	// 返回数据
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
