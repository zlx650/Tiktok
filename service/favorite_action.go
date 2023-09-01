package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
)

func FavoriteAction(videoId, actionType string) error {

	if actionType == "1" {
		err := dao.UpdateVideoIsFavoriteByVideoID(videoId, actionType)
		if err != nil {
			return fmt.Errorf("点赞失败: %s", err)
		}
	} else if actionType == "2" {
		err := dao.UpdateVideoIsFavoriteByVideoID(videoId, actionType)
		if err != nil {
			return fmt.Errorf("取消点赞失败: %s", err)
		}
	}

	return errors.New("UnknownError")

}
