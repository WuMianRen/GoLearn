/*
@Time : 2020/9/9 13:55
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

type Options struct {
	Interval time.Duration
}

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1})
	go reqTask(vCtx, "work1")
	go reqTask(vCtx, "work2")

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)
}
