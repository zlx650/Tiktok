package util

import (
	"fmt"
	"strconv"
	"time"
)

const IP = "127.0.0.1"
const Port = "8080"

var supportedFormats = map[string]struct{}{
	".mp4": {},
	//	todo 添加新支持的视频格式 ~
}

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

func GetDataUrl(name string) string {
	url := fmt.Sprintf("http://%s:%d/static/%s", IP, Port, name)
	return url
}

func IsSupportedVideoFormat(ext string) bool {
	_, ok := supportedFormats[ext]
	return ok
}

func GenerateUniqueFileName(userid int64, ext string) string {
	timestamp := time.Now().Unix() / int64(time.Millisecond)
	return fmt.Sprintf("%d_%d%s", userid, timestamp, ext)
}

func GenerateThumbnail(videoPath string, thumbnailPath string) (string, error) {
	//	todo 生成视频封面
	return videoPath + thumbnailPath, nil
}
