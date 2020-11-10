package main

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Push("3w132")
	fmt.Println(q.Pop())
}
