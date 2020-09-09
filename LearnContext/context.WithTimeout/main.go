/*
@Time : 2020/9/9 14:01
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
			fmt.Println("send request", name)
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTask(ctx, "work1")
	go reqTask(ctx, "work2")

	time.Sleep(time.Second * 3)

	// 这样写在 cancel 之前 协程已经停止
	fmt.Println("before cancel")
	cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("end")

}
