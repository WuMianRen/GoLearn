package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "woca")
	})

	r.GET("/user/:name", func(c *gin.Context) {

		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})



	r.Run()
}
