package service

import (
	"tiktok/dao"
	"tiktok/util"
	"log"
)

func GetPublishList(userID int64) ([]util.Video, error) {
	log.Println("service publish sucess")
	// 获取用户发布的视频列表
	videoList, err := dao.GetVideosByUserID(userID)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
