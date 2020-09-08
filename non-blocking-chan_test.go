package main

import (
	"fmt"
	"testing"
	"time"
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

// 非阻塞式的channel队列 ==> 傻瓜式地实现了先进后出队列
func TestPush1(t *testing.T) {
	q := make(chan int, 5)
	x := 100

	go func() {
		for i := 1; i < x; i++ {
			select {
			case q <- i:
				fmt.Println("PRODUCER:insert into channel: ", i)
			default:
				item := <-q
				fmt.Println("PRODUCER:channel is full,this value will be discard", item)
				q <- i
				fmt.Println("PRODUCER:insert into channel: ", i)
			}
			time.Sleep(time.Millisecond * 150)
		}
	}()

	// 消费者
	for i := 0; i < cap(q); i++ {
		go func() {
			for {
				select {
				case item := <-q:
					fmt.Println("CONSUMER:item :", item)
				default:
					fmt.Println("CONSUMER:queue empty")
				}
				time.Sleep(time.Second)
			}

		}()
	}

	time.Sleep(time.Hour)
}

// 非阻塞式的channel队列 ==> 但是把最新的消息丢弃了
func TestPush2(t *testing.T) {
	q := make(chan int, 5)
	x := 100
	go func() {
		for i := 1; i < x; i++ {
			select {
			case q <- i:
				fmt.Println("insert into channel: ", i)
			default:
				fmt.Println("channel is full,this value will be discard", i)
			}
			time.Sleep(time.Second)
		}
	}()

	// 消费者
	go func() {
		for {
			select {
			case item := <-q:
				fmt.Println("item :", item)
			default:
				fmt.Println("queue empty")
			}
			time.Sleep(2 * time.Second)
		}

	}()

	time.Sleep(time.Hour)
}
