/*
@Time : 2020/9/8 10:48
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 热加载
	// go get -v -u github.com/pilu/fresh
	// go run main.go命令换成fresh
	r := gin.Default()

	r.GET("/1", func(c *gin.Context) {
		c.String(http.StatusOK, "1")
	})

	r.GET("/2", func(c *gin.Context) {
		c.String(http.StatusOK, "1")
	})
}
