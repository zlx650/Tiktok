package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
	"tiktok/service"
	"tiktok/util"
)

func PublishVideoController(context *gin.Context) {
	//验证token后从context获取的user_id
	uidRaw, _ := context.Get("user_id")
	uid, ok := uidRaw.(int64)
	if !ok {
		PublishVideoControllerErrorResponse(context, "user_id解析出错")
		return
	}
	title := context.PostForm("title")
	form, err := context.MultipartForm()
	if err != nil {
		PublishVideoControllerErrorResponse(context, err.Error())
		return
	}
	files := form.File["data"]
	for _, file := range files {
		//	check 是否为支持的视频格式
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !util.IsSupportedVideoFormat(ext) {
			PublishVideoControllerErrorResponse(context, "不支持上传该视频格式")
			return
		}

		//  生成唯一的文件名用于保存
		filename := util.GenerateUniqueFileName(uid, ext)

		//	写入static
		savePath := "./static/video" + filename
		err = context.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoControllerErrorResponse(context, err.Error())
			return
		}
		//  todo 加一个视频封面
		coverName := ""

		//  数据持久化
		err = service.PublishVideo(uid, filename, coverName, title)
		if err != nil {
			PublishVideoControllerErrorResponse(context, err.Error())
		}
		PublishVideoControllerSuccessResponse(context, file.Filename+"上传成功")
	}

}

func PublishVideoControllerErrorResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: msg})
}

func PublishVideoControllerSuccessResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: msg})
}
