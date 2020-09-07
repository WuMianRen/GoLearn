/*
@Time : 2020/9/7 15:08
@Author : wumian
@File : string
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

// 字符串使用 UTF8 编码，UTF8 的好处在于，如果基本是英文，每个字符占 1 byte， 和 ASCII 编码是一样的，非常节省空间，
// 如果是中文，一般占3字节。包含中文的字符串的处理方式与纯 ASCII 码构成的字符串有点区别。
func main() {
	str1 := "Golang"
	str2 := "Go语言" //每个中文字符占三个
	// Kind 返回当前绑定的类型
	// reflect.TypeOf().Kind()  可以知道某个变量的类型
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) //uint8  字符串是以 byte 数组形式保存的
	fmt.Println(str1[2], string(str1[2]))       //108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])     //232 è
	fmt.Println("len(str2)", len(str2))         //8

	// 正确的处理方法 --  string 转为 rune 数组
	str3 := "Go语言" //每个中文字符占三个
	runeArr := []rune(str3)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
	fmt.Println("len(runeArr)", len(runeArr))      // 4

}
