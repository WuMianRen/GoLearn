/*
@Time : 2020/9/7 15:40
@Author : wumian
@File : map
@Software: GoLand
*/
package main

import "fmt"


// java 当中是HashMap Python C# Dictionary 中是 dict 使用方式差不多
var m1 = make(map[string]int)

func main() {
	// 声明赋值
	a2 := map[string]string{
		"age":  "1",
		"name": "nihao",
	}

	m1["Tom"] = 18

	fmt.Println(a2, m1)
}
