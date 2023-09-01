package dao

import (
	"gorm.io/gorm"
	"sync"
	"tiktok/config"
	"tiktok/models"
	"time"
)

type VideoDao struct{}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) CreateVideo(video *models.Video) error {
	if err := DB.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (*VideoDao) QueryVideoCountByUserId(userId int64, count *int64) error {
	if err := DB.Model(&models.Video{}).Where("author_id = ?", userId).Count(count).Error; err != nil {
		return err
	}
	return nil
}

func QueryFeedVideoList(postTime time.Time) ([]models.Video, error) {
	// 从数据库中取videoList数据
	var videoList []models.Video
	if err := DB.Preload("Author").Where("post_time < ?", postTime).Order("post_time desc").Limit(30).Find(&videoList).Error; err != nil {
		return nil, err
	}

	return videoList, nil
}

func UpdateVideoURL() error {

	newUrl := config.IP

	if err := DB.Model(&models.Video{}).
		Where("play_url LIKE ?", "%static%").
		Update("play_url", gorm.Expr("REGEXP_REPLACE(play_url, ?, ?)", "^https:.*\\.com", newUrl)).Error; err != nil {
		return err
	}

	return nil
}
