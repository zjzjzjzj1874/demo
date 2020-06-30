package main

import (
	"flag"
	"fmt"
)

// flag包的运行方式:(remember:在运行时记得执行flag.Parse()方法!!!)
// 1.go run main.go -greet=goodbye -debug=true
// 2.go build后,./main -debug=true -greet=goodbye
// 3.编辑goland的configuration中心,添加对应的启动参数
// go run main.go -h ==> 可以查看哪些启动参数可以配置
// 更多玩法,敬请参考:https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter13/13.1.html
func flagUse() {
	d := new(bool)
	flag.BoolVar(d, "debug", false, "Whether debug mode is enabled")
	v := flag.Bool("v", false, "show version")
	flag.Parse()
	if *v {
		fmt.Println("v1.0.0")
	}
	if *d {
		fmt.Println(*d)
	}
}
