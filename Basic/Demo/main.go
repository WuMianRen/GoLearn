package main

import (
	"fmt"
	
	"example/calc"
	"example/incr"
)

func main() {
	incr.Add2()
	fmt.Println(calc.Add(10, 3))
}
