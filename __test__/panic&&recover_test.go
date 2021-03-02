// golang中常见的panic有空指针,切片数组下表越界,断言错误,被除数为0,操作关闭的通道,
package __test__

import (
	"fmt"
	"runtime"
	"testing"
)

type s struct {
	name string
	age  int
}

// panic保护的函数
func protectRun(f func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error: // 运行时错误
				fmt.Println("运行时错误 error:", err)
			default: // 非运行时错误
				fmt.Println("非运行时错误 error:", err)
			}
		}
	}()
	f()
}

// 1.测试panic是否可恢复 -- nil pointer ==> 可以恢复;
func TestNilPointerPanicWithRecover(t *testing.T) {
	fmt.Println("运行前")

	// NilPointer 空指针异常的panic可恢复
	protectRun(func() {
		fmt.Println("有recover的panic前")
		var ss *s
		ss.name = "hello"
		fmt.Printf("有recover的panic后")
	})

	// 以下的会导致程序退出
	func() {
		fmt.Println("无recover的panic前")
		var ss *s
		ss.name = "hello"
		fmt.Printf("无recover的panic后")
	}()

	fmt.Println("运行结束")
}

// 2.测试panic是否可恢复 -- index out of range ==> 可以恢复
func TestPanicWithRecover(t *testing.T) {
	fmt.Println("运行前")

	protectRun(func() {
		fmt.Println("有recover的panic前")
		ss := make([]int, 1)
		ss[5] = 5
		fmt.Printf("有recover的panic后")
	})

	// 以下的会导致程序退出
	func() {
		fmt.Println("无recover的panic前")
		ss := make([]int, 1)
		ss[5] = 5
		fmt.Printf("无recover的panic后")
	}()

	fmt.Println("运行结束")
}

// 3.测试panic是否可恢复 -- type assert ==> 可以恢复;
func TestPanicWithTypeAssert(t *testing.T) {
	fmt.Println("运行前")

	//protectRun(func() {
	//	fmt.Println("有recover的panic前")
	//	var ss interface{}
	//	ss = "hello world"
	//	fmt.Println(ss.(int))
	//	fmt.Printf("有recover的panic后")
	//})

	// 以下的会导致程序退出
	func() {
		fmt.Println("无recover的panic前")
		var ss interface{}
		ss = "hello world"
		fmt.Println(ss.(int))
		fmt.Printf("无recover的panic后")
	}()

	fmt.Println("运行结束")
}

// 4.测试panic是否可恢复 -- divide zero ==> 可以恢复;
func TestPanicWithZeroDivide(t *testing.T) {
	fmt.Println("运行前")

	protectRun(func() {
		fmt.Println("有recover的panic前")
		var ss int
		fmt.Println(10 / ss)
		fmt.Printf("有recover的panic后")
	})

	// 以下的会导致程序退出
	func() {
		fmt.Println("无recover的panic前")
		var ss int
		fmt.Println(10 / ss)
		fmt.Printf("无recover的panic后")
	}()

	fmt.Println("运行结束")
}

// 往nil的channel中写入数据
// 5.测试panic是否可恢复 -- nil channel ==> 无法恢复 报错=> fatal error: all goroutines are asleep - deadlock!;
func TestPanicWithNilChan(t *testing.T) {
	fmt.Println("运行前")

	protectRun(func() {
		fmt.Println("有recover的panic前")
		var ss chan int
		ss <- 10
		fmt.Printf("有recover的panic后 \n")
	})

	fmt.Println("运行结束")
}

// 往长度为0的channel中写入数据
// 6.测试panic是否可恢复 -- zero channel ==> 无法恢复 报错=> fatal error: all goroutines are asleep - deadlock!;
func TestPanicWithZeroChan(t *testing.T) {
	fmt.Println("运行前")

	protectRun(func() {
		fmt.Println("有recover的panic前")
		ss := make(chan int) // make(chan int,0) 后面注释里的写法没有panic
		ss <- 10
		fmt.Printf("有recover的panic后 \n")
	})

	fmt.Println("运行结束")
}

// 往closed的channel中写入数据
// 7.测试panic是否可恢复 -- closed channel ==> 无法恢复 报错=> fatal error: all goroutines are asleep - deadlock!;
func TestPanicWithClosedChan(t *testing.T) {
	fmt.Println("运行前")

	protectRun(func() {
		fmt.Println("有recover的panic前")
		ss := make(chan int,1) // make(chan int,0) 后面注释里的写法没有panic
i
		ss <- 10
		fmt.Printf("有recover的panic后 \n")
	})

	fmt.Println("运行结束")
}

// fatal err和runtime err,为什么fatal err不可恢复

// 测试copy的
func TestCopy(t *testing.T) {
	a := []s{{
		name: "chiva",
		age:  18,
	}}
	ss := make([]s, len(a))
	fmt.Println("before a = ", a)
	copy(ss, a)
	for idx := range ss {
		//ss[idx].change()
		changes(&ss[idx])
	}
	//
	//changes(&ss[0])

	fmt.Println("after a = ", a)
	fmt.Println("ss == ", ss)
}

func (s *s) change() {
	s.name = "chiva1"
	s.age = 17
}
func changes(s2 *s) *s {
	s2.name = "chiva2"
	s2.age = 16
	return s2
}
