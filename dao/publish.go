package dao

import (
	"log"
	"tiktok/models"
)

func GetVideosByUserID(userID int64) ([]models.Video, error) {
	var videoList []models.Video
	log.Println("dao publish sucess")
	err := DB.Where("user_id = ?", userID).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	log.Println("视频发布列表:", videoList)
	return videoList, nil
}
