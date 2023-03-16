package main

import (
	"github.com/GoFeGroup/gordp/glog"
	"github.com/gin-gonic/gin"
)

func main() {
	glog.SetLevel(glog.NONE)

	r := gin.Default()
	r.GET("/api/rdp", rdpProxy)
	r.Use(feMw("/"))
	_ = r.Run(":8081")
}
