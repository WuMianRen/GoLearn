/*
@Time : 2020/9/7 15:45
@Author : wumian
@File : pointer
@Software: GoLand
*/
package main

import "fmt"

func main() {
	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"         // 这里直接使用指针修改值
	fmt.Println(str)     // Hello 修改了p  所以 str 也需要变

	// 一般 指针通常用于在参数中传递  或者给某个类型定义新方法的时候使用
	// Go 当中  参数是按值传递的 如果不使用指针 函数内部会拷贝一份参数的副本
	// 对参数的修改不会影响外部变量的值 如果参数使用指针 则会印象外部的变量

	num := 100
	add(num)
	fmt.Println(num) // num 无变化
	realAdd(&num)    // num 加 1
	fmt.Println(num)

}

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}
