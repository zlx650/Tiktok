package dao

import (
	"tiktok/models"
	"tiktok/util"
)

func UpdateVideoIsFavoriteByVideoID(videoId, actionType string) error {

	if err := util.DB.Model(&models.Video{}).Where("video_id = ?", videoId).Update("is_favorite", actionType).Error; err != nil {
		return err
	}

	return nil
}
