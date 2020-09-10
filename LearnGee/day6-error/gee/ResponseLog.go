/*
@Time : 2020/9/10 13:52
@Author : wumian
@File : ccc
@Software: GoLand
*/
package gee

import (
	"bytes"
	"fmt"
	"strings"
)

func CommonLogInterceptor() HandlerFunc {
	return func(context *Context) {
		strBody := ""
		blw := bodyLogWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: context.Writer,
		}
		context.Writer = blw
		context.Next()

		strBody = strings.Trim(blw.bodyBuf.String(), "\n")
		if len(strBody) > MaxPrintBodyLen {
			strBody = strBody[:(MaxPrintBodyLen - 1)]
		}
		fmt.Println("end req[" + context.Req.RequestURI + "], res[" + strBody + "]")
	}
}
