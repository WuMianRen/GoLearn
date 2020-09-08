/*
@Time : 2020/9/8 10:41
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	r := gin.Default()

	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Logger())

	// 作用于单个
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 作用于组
	//auth := r.Group("/")
	//
	//auth.Use((AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//}

	// 自定义中间件

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context实例设置一个值
		c.Set("geektutu", "1111")

		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}
