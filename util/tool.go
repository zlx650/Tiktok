package util

import (
	"strconv"
	"time"
)

// ConvertTimeToTimestampStr 接收一个time.Time的类型数据，返回string类型、以s为单位的时间戳
func ConvertTimeToTimestampStr(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

func ConvertTimestampStrToUnix(timestampStr string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64) // 将时间戳字符串解析为 int64 类型的时间戳值
	t := time.Unix(timestamp, 0)                             // 使用 time.Unix() 函数将时间戳值转换为 time.Time 类型
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
