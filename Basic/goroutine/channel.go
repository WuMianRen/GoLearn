/*
@Time : 2020/9/7 16:54
@Author : wumian
@File : channel
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

// 使用 channel 信道，可以在协程之间传递消息。阻塞等待并发协程返回消息。
var ch = make(chan string, 10)

func main() {
	for i := 0; i < 3; i++ {
		go download2("a.com/" + string(i+'0'))
	}

	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finsh", msg)
	}

	fmt.Println("Done")
}

func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url
}
