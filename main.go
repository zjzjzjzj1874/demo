package main

import (
	"fmt"
	"math"
)

func main() {
	flagUse()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(i)
	//	}
	//}()

	//select {} // select阻塞,可以确保go协成执行完毕 ==> 但是前面的执行完之后,就会进入死锁阶段.

	//sb := strings.Builder{}
	//sb.WriteString("hello")
	//sb.WriteString("world")
	//str := sb.String()
	//fmt.Println(str, sb.Cap(), sb.Len())
	//
	//src := "abcdefg123abc"
	//sub := "abcg"
	//fmt.Println(strings.LastIndex(src, sub))
	//fmt.Println(strings.Contains(src, sub))
	//fmt.Println(strings.Count(src, sub))

	fmt.Println(math.Pow(2.0, 3.0))

}
