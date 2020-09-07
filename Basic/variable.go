/*
@Time : 2020/9/7 15:02
@Author : wumian
@File : variable.go
@Software: GoLand
*/
package main

import "fmt"

var a1 int     // 如果没有赋值，默认为0
var a2 int = 1 // 声明时赋值
var a3 = 1     // 声明时赋值

func main() {
	// 简单类型
	// 1. nil
	// 2. int (8 16 32 64 ) uint(8 16 )..
	// 3. float32 float 64
	// 4. byte = uint8
	// 5. string 字符串使用 utf8
	// 6. boolean
	msg := "Hello world"
	fmt.Println(msg)
}
