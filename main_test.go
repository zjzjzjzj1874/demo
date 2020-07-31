package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
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

// 浮点数与字符串转化
func TestStringParseToFloat64(t *testing.T) {
	tagVal, err := strconv.ParseUint("9765439183791093097", 10, 64)
	fmt.Println(tagVal, err)
}

// 不用waitGroup测试
func TestWithoutWaitGroup(t *testing.T) {
	intArr := make([]int, 20)
	var param string
	for idx := range intArr {
		param += fmt.Sprintf("%d,", idx)
	}
	postHttp(param)
}

// waitGroup用法测试
func TestWaitGroup(t *testing.T) {
	intArr := make([]int, 20)
	wg := sync.WaitGroup{}

	var (
		param string
		go0   string
		go1   string
		go2   string
		go3   string
	)
	for idx := range intArr {
		switch idx % 4 {
		case 0:
			go0 += fmt.Sprintf("%d,", idx)
		case 1:
			go1 += fmt.Sprintf("%d,", idx)
		case 2:
			go2 += fmt.Sprintf("%d,", idx)
		default:
			go3 += fmt.Sprintf("%d,", idx)
		}
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		switch i {
		case 0:
			param = go0
		case 1:
			param = go1
		case 2:
			param = go2
		case 3:
			param = go3
		}
		go postHttpWithWaitGroup(param, &wg)
	}

	wg.Wait()
	fmt.Println("all work has done!")
}

func postHttpWithWaitGroup(param string, group *sync.WaitGroup) {
	defer group.Done()
	resp, err := http.Get("http://localhost:8080/hello?param=" + param)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	fmt.Println(string(content))
}

func postHttp(param string) {
	resp, err := http.Get("http://localhost:8080/hello?param=" + param)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	fmt.Println("postHttp : ", string(content))
}
