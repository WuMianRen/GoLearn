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

const MaxPrintBodyLen = 512

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

		if len(strBody) > MaxPrintBodyLen {
			strBody = strBody[:(MaxPrintBodyLen - 1)]
		}

		log.Printf("[%d] %s %s in %v",
			c.StatusCode, c.Req.RequestURI, strBody, time.Since(t))
	}
}
