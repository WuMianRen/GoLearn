/*
@Time : 2020/9/7 16:30
@Author : wumian
@File : struct
@Software: GoLand
*/
package main

import "fmt"

//
type Student struct {
	name string
	age  string
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s ,i am %s", person, stu.name)
}

func main() {
	// 实例化方式一
	stu := &Student{name: "xiaoming", age: "18"}
	msg := stu.hello("Jack")
	fmt.Println(msg)
	// 实例化方式二
	stu2 := new(Student)
	fmt.Println(stu2.hello("Alice"))

	// interfaces
	var p Person = &Student{
		name: "Tom",
		age:  "18",
	}
	fmt.Println(p.getName()) // Tom

	stu3 := p.(*Student) // 将接口强制转换为示例
	fmt.Println(stu3.getName())

}

// interfaces
// 并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可。
type Person interface {
	getName() string
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

// 这样的方式可以检测到是否实现了接口所有的方法
var _ Person = (*Student)(nil)
var _ Person = (*Worker)(nil)

// 空接口

func main2() {
	m := make(map[string]interface{})
	m["1"] = "1"
	m["2"] = 1
	m["3"] = [3]int{1, 2, 3}
}
