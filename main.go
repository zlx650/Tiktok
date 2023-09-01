package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/routers"
)

func main() {
	r := gin.Default()

	routers.BaseRoutersInit(r)
	routers.InteractRoutersInit(r)

	r.Run()

}
