package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	r.GET("/index", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"垃圾":  "是的",
			"垃2圾": "是的asd",
		})
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
