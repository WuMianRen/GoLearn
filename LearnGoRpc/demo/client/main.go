/*
@Time : 2020/9/8 11:27
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {

	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	//// 同步方法
	//if err := client.Call("Cal.Square", 12, &result); err != nil {
	//	log.Fatal("Failed to call Cal.Square. ", err)
	//}

	// 异步方法
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
