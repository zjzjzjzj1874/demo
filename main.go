package main

import (
	"flag"
	"fmt"
)

// flag包的运行方式:(remember:在运行时记得执行flag.Parse()方法!!!)
// 1.go run main.go -greet=goodbye -debug=true
// 2.go build后,./main -debug=true -greet=goodbye
// 3.编辑goland的configuration中心,添加对应的启动参数
var (
	filePath = flag.Bool("debug", false, "")
	str      = flag.String("greet", "hello", "")
)

func main() {
	flag.Parse()
	fmt.Println(*filePath)
	fmt.Println(*str)
}
