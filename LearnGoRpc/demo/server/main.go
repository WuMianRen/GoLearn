/*
@Time : 2020/9/8 11:24
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//使用 Golang 的标准库 net/rpc
//func (t *T) MethodName(argType T1, replyType *T2) error

// 1) 方法类型（T）是导出的（首字母大写）
// 2) 方法名（MethodName）是导出的
// 3) 方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
// 4) 方法的第2个参数一个指针(replyType *T2)
// 5) 方法的返回值类型是 error

// example

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	err := rpc.Register(new(Cal))
	if err != nil {
		log.Fatal("Error rpcRegister: ", err)
	}
	rpc.HandleHTTP()

	log.Printf("Server Rpc server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
