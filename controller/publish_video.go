package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/service"
)

func PublishVideoController(context *gin.Context) {
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
		//	check 是否为视频
		//文件名 userid +
		filename := "test"
		//	保存于static
		savePath := "./static" + filename
		err = context.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoControllerErrorResponse(context, err.Error())
			return
		}
		err = service.PublishVideo(uid, filename, "", title)
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
