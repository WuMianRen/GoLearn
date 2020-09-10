/*
@Time : 2020/9/10 10:56
@Author : wumian
@File : logger
@Software: GoLand
*/
package gee

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"time"
)

const MAX_PRINT_BODY_LEN = 512

type bodyLogWriter struct {
	http.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		strBody := ""
		blw := bodyLogWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		c.Next()
		strBody = strings.Trim(blw.bodyBuf.String(), "\n")

		if len(strBody) > MAX_PRINT_BODY_LEN {
			strBody = strBody[:(MAX_PRINT_BODY_LEN - 1)]
		}

		log.Printf("[%d] %s body:%s in %v",
			c.StatusCode, c.Req.RequestURI, strBody, time.Since(t))
	}
}
