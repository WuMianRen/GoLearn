/*
@Time : 2020/9/7 15:54
@Author : wumian
@File : if_for_switch
@Software: GoLand
*/
package main

import "fmt"

func main() {

	// if
	age := 18
	if age < 18 {
		fmt.Println("kid")
	} else {
		fmt.Println("Adult")
	}

	// 简写
	if age := 18; age < 18 {
		fmt.Println("kid")

	} else {
		fmt.Println("Adult")
	}

	//switch
	fmt.Println("------switch-----")
	type Gender int8

	// go 当中没有枚举  这里可以使用常量来替代枚举
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	// 在 go 当中的  switch 不需要 break
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough // 如果需要 switch 继续向下走就需要这个
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}

	// for
	fmt.Println("--------for--------")
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}

	// 对 arr slice map 可以使用 for range
	nums := []int{1, 2, 3, 4, 5}

	for i, num := range nums {
		fmt.Println(i, num)
	}

	nums = append(nums, 10, 11, 12)
	for i, num := range nums {
		fmt.Println(i, num)
	}

	m2 := map[string]string{
		"name": "zhagnsan",
		"age":  "20",
	}
	for s, s2 := range m2 {
		fmt.Println(s, s2)
	}
}
