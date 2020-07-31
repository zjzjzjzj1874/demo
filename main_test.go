package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
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

// IOT gateway测试
func TestIotGateway(t *testing.T) {
	// function: 生成IOT gateway需要的鉴权token
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authBytes := md5.Sum([]byte("HXBYcA32PgQhiHAsQB" + timestamp))
	fmt.Println(fmt.Sprintf("%s;%x;%s", "rockontrol", authBytes, timestamp))

}
