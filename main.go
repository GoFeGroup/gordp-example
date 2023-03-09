package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/rdp", rdpProxy)
	r.Use(feMw("/"))
	_ = r.Run(":8081")
}
