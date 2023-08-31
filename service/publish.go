package service

import (
	"log"
	"tiktok/dao"
	"tiktok/models"
)

func GetPublishList(userID int64) ([]models.Video, error) {
	log.Println("service publish sucess")
	// 获取用户发布的视频列表
	videoList, err := dao.GetVideosByUserID(userID)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
