package main

import (
	"errors"
	"fmt"
)

var (
	Err1 = errors.New("err1")
	Err2 = errors.New("err2")
)

// how err working
func MyTestErr() {
	//e1 := errors.New("err1")
	e1 := Err1
	fmt.Println("e1 === ", &e1, "err1 ==== ", &Err1)
	if e1 == Err1 {
		fmt.Println("e1 equal Err1")
	} else {
		fmt.Println("e1 not equal Err1")
	}

	if e1.Error() == Err1.Error() {
		fmt.Println("e1.Error() equal Err1.Error()")
	} else {
		fmt.Println("e1.Error() not equal Err1.Error()")
	}
}
