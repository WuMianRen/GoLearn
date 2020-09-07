package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	que, rem := div(100, 17)
	fmt.Println(que, rem)
	fmt.Println(add(100, 100))

	errorShow()
	if err := customErr(""); err != nil {
		fmt.Println(err)
	}

	// get(5)
	getAdv(5)
}

// 参数 返回值
func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 * num2
}

// 错误处理
func errorShow() {
	// os.Open 返回两个值
	// *File error
	_, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
}

//  errors.New 定义错误
func customErr(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	return nil
}

// 程序非正常退出  panic

func get(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}

// 替代 try catch 的机制   defer recover
func getAdv(index int) (ret int) {
	// TODO   可深入理解  defer 机制
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened", r)
			ret = -1
		}
	}()
	arr := [3]int{1, 2, 3}
	return arr[index]
}
