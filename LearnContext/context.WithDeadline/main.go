/*
@Time : 2020/9/9 14:08
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
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(time.Second * 1)
		}
	}

}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))

	go reqTask(ctx, "work1")
	go reqTask(ctx, "work2")

	// 这这里将协程的死亡时间执行为当前一秒后
	// 所以在执行 cancel 之前 协程已经死亡
	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(2 * time.Second)
}
