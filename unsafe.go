package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// 定义一个和 strings 包中的 Reader 相同的本地结构体
type Reader struct {
	s        string
	i        int64
	prevRune int
}

// 类型转换 ==> 将strings的Reader对象转为本地的Reader,然后对strings的Reader进行修改
func TypeConversion() {
	// 创建一个 strings 包中的 Reader 对象
	sr := strings.NewReader("abcdef")
	// 此时 sr 中的成员是无法修改的
	fmt.Println("init value:", sr)

	// 我们可以通过 unsafe 来进行修改
	// 先将其转换为通用指针
	p := unsafe.Pointer(sr)
	// 再转换为本地 Reader 结构体
	pR := (*Reader)(p)
	// 这样就可以自由修改 sr 中的私有成员了
	(*pR).i = 3              // 修改索引
	(*pR).s = "hello world!" // 修改索引
	// 看看修改结果
	fmt.Println("after modify:", sr)
	// 看看读出的是什么
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
	fmt.Println("after modify:", sr)
}

func PointerRange() {
	// 创建一个 strings 包中的 Reader 对象
	// 它有三个私有字段：s string、i int64、prevRune int
	sr := strings.NewReader("hello unsafe pointer!")
	// 此时 sr 中的成员是无法修改的
	fmt.Println("before modify:", sr)
	// 但是我们可以通过 unsafe 来进行修改
	// 先将其转换为通用指针
	p := unsafe.Pointer(sr)
	// 获取结构体地址
	up0 := uintptr(p)
	// 确定要修改的字段（这里不能用 unsafe.Offset of 获取偏移量，因为是私有字段）
	if sf, ok := reflect.TypeOf(*sr).FieldByName("i"); ok {
		// 偏移到指定字段的地址
		up := up0 + sf.Offset
		// 转换为通用指针
		p = unsafe.Pointer(up)
		// 转换为相应类型的指针
		pi := (*int64)(p)
		// 对指针所指向的内容进行修改
		*pi = 3 // 修改索引
	}

	if sf, ok := reflect.TypeOf(*sr).FieldByName("s"); ok {
		// 偏移到指定字段的地址
		up := up0 + sf.Offset
		// 转换为通用指针
		p = unsafe.Pointer(up)
		// 转换为相应类型的指针
		pi := (*string)(p)
		// 对指针所指向的内容进行修改
		*pi = "hello" // 修改索引
	}
	// 看看修改结果
	fmt.Println("after modify:", sr)
	// 看看读出的是什么
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}

type Foo struct {
	A int
	B int
}

func ModifyArray() {
	foo := &Foo{1, 2}
	fmt.Println(foo)

	base := uintptr(unsafe.Pointer(foo))
	offset := unsafe.Offsetof(foo.A)
	//offset := unsafe.Offsetof(foo.B)

	ptr := unsafe.Pointer(base + offset)
	*(*int)(ptr) = 3

	fmt.Println(foo)
}
