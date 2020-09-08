/*
@Time : 2020/9/8 11:00
@Author : wumian
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &Student{
		Name:   "zhagnsan",
		Male:   true,
		Scores: []int32{100, 101, 102},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error", err)
	}

	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)

	if err != nil {
		log.Fatal("unmarshaling ", err)
	}
	fmt.Println(test.GetName())
	fmt.Println(newTest.GetName())
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

}
