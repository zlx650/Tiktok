package util

import "time"

type Video struct {
}

// TableName 表示配置操作数据库的表名称
func (Video) TableName() string {
	return "video"
}
