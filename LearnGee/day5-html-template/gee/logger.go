/*
@Time : 2020/9/10 10:56
@Author : wumian
@File : logger
@Software: GoLand
*/
package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
