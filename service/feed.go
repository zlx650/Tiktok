package service

import (
	"gorm.io/gorm"
	"tiktok/models"
	"tiktok/util"
	"time"
)

func QueryFeedVideo(postTime time.Time) []models.Video {

	// 从数据库中取videoList数据
	var videoList []models.Video

	util.DB.Preload("Author").
		Where("post_time < ?", postTime).
		Order("post_time desc").
		Limit(30).
		Find(&videoList)

	return videoList
}

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

func UpdateVideoURL() error {

	newUrl := util.IP

	if err := util.DB.Model(&models.Video{}).
		Where("play_url LIKE ?", "%static%").
		Update("play_url", gorm.Expr("REGEXP_REPLACE(play_url, ?, ?)", "^https:.*\\.com", newUrl)).Error; err != nil {
		return err
	}

	return nil
}
