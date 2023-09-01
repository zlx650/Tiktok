package service

import (
	"tiktok/models"
	"time"
)

func FindEarliestPostTime(videoList []models.Video) int64 {
	var nextTime int64 = time.Now().Unix()
	if len(videoList) > 0 {
		for _, video := range videoList {
			videoTime := video.PostTime.Unix()
			if videoTime < nextTime {
				nextTime = video.PostTime.Unix()
			}
		}
	}
	return nextTime
}
