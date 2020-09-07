/*
@Time : 2020/9/7 15:19
@Author : wumian
@File : array_slice
@Software: GoLand
*/
package main

import "fmt"

// 声明数组
var arr [5]int     // 一维
var arr2 [5][5]int //二维

// 声明初始化
var arr3 = [5]int{1, 2, 3, 4, 5}

// or arr4 := [5]int{1, 2, 3, 4, 5}

func main() {
	// 数组 需要声明长度
	arr5 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr5); i++ {
		arr5[i] += 100
	}

	fmt.Println(arr5)

	// 切片是数组的抽象 切边使用数组未底层结构
	// 切片包含三个组件 1. 容量 2. 长度 3. 指向底层数组的指针
	// 切片可以随时进行扩展

	_ = make([]float32, 0)          //长度为 0 的切片
	slice2 := make([]float32, 3, 5) // [0,0,0] 长度为三 容量为5的切片
	fmt.Println(len(slice2), cap(slice2))

	// 切片容量可以根据需要扩展
	slice2 = append(slice2, 1, 2, 3, 4) // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2))

	fmt.Println("-------------")
	// 子切片 [start, end]  左开右闭
	sub1 := slice2[3:]  //[1, 2, 3, 4]
	sub2 := slice2[:3]  // [0, 0, 0]
	sub3 := slice2[1:4] //  [0 0 1]

	combined := append(sub1, sub2...)

	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
	fmt.Println(combined)
}
