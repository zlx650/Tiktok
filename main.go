package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"tiktok/dao"
	"tiktok/routers"
)

func main() {

	if err := Init(); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()

	routers.BaseRoutersInit(r)
	routers.InteractRoutersInit(r)

	err := r.Run()
	if err != nil {
		return
	}

}

func Init() error {
	if err := dao.Init(); err != nil {
		os.Exit(-1)
	}

	return nil
}
