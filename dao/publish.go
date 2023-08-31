package dao

import (
	"tiktok/util"
	"log"
)

func GetVideosByUserID(userID int64) ([]util.Video, error) {
	var videoList []util.Video
  log.Println("dao publish sucess")
	err := util.DB.Where("author_id = ?", userID).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	log.Println("视频发布列表:",videoList)
	return videoList, nil
}
