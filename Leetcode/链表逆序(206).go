/*
@Time : 2020/12/22 10:52
@Author : wumian
@File : 链表逆序(206)
@Software: GoLand
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Name string
	Next *ListNode
}

func main() {

	a := &ListNode{"1", nil}
	b := &ListNode{"2", nil}
	c := &ListNode{"3", nil}
	d := &ListNode{"4", nil}
	e := &ListNode{"5", nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	// reverseV1(a).listPrint()
	reverseV2(a).listPrint()

}

// 逆置法
func reverseV1(head *ListNode) *ListNode {

	var reverse *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = reverse
		reverse = head
		head = next
		fmt.Println("old", head)
		fmt.Println("reverse", reverse)
	}

	return reverse
}

// 头插法
func reverseV2(head *ListNode) *ListNode {

	temp := &ListNode{"0", nil}
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = temp.Next
		temp.Next = head
		head = next
		fmt.Println("old", head)
		fmt.Println("reverse", temp)
	}

	return temp.Next
}

func length(list *ListNode) (int, []*ListNode) {
	temp := make([]*ListNode, 5)
	count := 0
	for list != nil {
		temp[count] = list
		list = list.Next
		count++
	}

	return count, temp
}

func (list *ListNode) listPrint() {
	list2 := list
	for list2 != nil {
		fmt.Println(list2.Name)
		list2 = list2.Next
	}
}
