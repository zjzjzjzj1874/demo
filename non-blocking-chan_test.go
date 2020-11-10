package main

import (
	"fmt"
	"testing"
	"time"
)

// 非阻塞式的channel队列 ==> 傻瓜式地实现了先进后出队列
func TestPush1(t *testing.T) {
	q := make(chan int, 1)
	x := 100

	// 一个消费者在不间断产生数据 ==> 超出队列的长度,自己消费最先进入的一个,然后把本次产生的加入channel队列中
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

	// 配置了channel容量个消费者 ==> 5个
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

// 非阻塞式的channel队列 ==> 傻瓜式地实现了先进后出队列
func TestPush(t *testing.T) {
	q := make(chan int, 5)
	x := 100

	// 2个消费者在不间断产生数据 ==> 超出队列的长度,自己消费最先进入的一个,然后把本次产生的加入channel队列中
	for i := 0; i < 2; i++ {
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
	}

	// 配置了channel容量个消费者 ==> 5个
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

// 阻塞式的channel队列 ==> 超出channel队列的长度后,将阻塞等待消费结束
func TestPush3(t *testing.T) {
	q := make(chan int, 5)
	x := 100
	go func() {
		for i := 1; i < x; i++ {
			select {
			case q <- i:
				fmt.Println("insert into channel: ", i)
				// 没有default分支,则该channel产生会阻塞
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
