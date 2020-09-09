/*
@Time : 2020/9/8 11:45
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	//done := make(chan int, 0)
	//js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	//<-done

	done := make(chan int, 0)
	btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc))
	<-done
}

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans")
	btnEle   = js.Global().Get("btn")
)

func fibFunc(this js.Value, args []js.Value) interface{} {
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err == nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}
