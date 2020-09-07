/*
@Time : 2020/9/7 17:01
@Author : wumian
@File : calc_test.go
@Software: GoLand
*/
package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}
