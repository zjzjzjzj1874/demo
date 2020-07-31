package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// 随机生成0-1的小数
func TestMyTestRand01(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(float64(rand.Intn(100)) / 100)
	}
}

// 切片测试
func TestMyTestSlice(t *testing.T) {
	arr := make([]int, 10)

	fmt.Println(arr[0:4])
	fmt.Println(arr[0:5])
}
