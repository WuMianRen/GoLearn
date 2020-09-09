/*
@Time : 2020/9/9 13:49
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 控制多个写协程
	ctx, cancel := context.WithCancel(context.Background())

	go reqTask(ctx, "work1")
	go reqTask(ctx, "work2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
