package main

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	q := make(chan int, 5)
	x := []int{1, 2, 3, 4, 5, 6}
	for _, value := range x {
		err := Push(q, value)
		fmt.Printf("error:%v\n", err)
	}

	for _, value := range x {
		fmt.Println(value)
		v, err := Get(q)
		fmt.Printf("v:%v, error:%v\n", v, err)
	}
}
