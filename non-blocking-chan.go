package main

import (
	"errors"
)

func Push(q chan int, item int) error {
	select {
	case q <- item:
		return nil
	default:
		return errors.New("queue full")
	}

}

// 这里可以用go for起的一个协成来处理,一直监听这个
func Get(q chan int) (int, error) {
	var item int
	select {
	case item = <-q:
		return item, nil
	default:
		return 0, errors.New("queue empty")
	}
}
