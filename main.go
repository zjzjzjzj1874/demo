package main

import (
	"demo/task"
	"fmt"
	"time"
)

func main() {
	// 位运算解决签到方案(每个月最多31天,in32就够用了,然后签到一天,将当天的位设置为1即可)
	const (
		a = 1 << iota // 1 << 0
		b             // 1 << 1
		c             // 1 << 2
		d             // 1 << 3
		e             // 1 << 4
		f             // 1 << 5
	)

	signIn := 0
	signIn |= a
	signIn |= d
	//fmt.Println(a, d, b, c, d, e, f, 1<<1, 3<<2)
	fmt.Println(" sign in:", signIn, signIn&a == 0, signIn&b == 0, signIn&c == 0, signIn&d == 0, signIn&e == 0, signIn&f == 0)

	//flagUse()
	//go func() {
	//	for i := 0; i < 10; i {
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

	//go task()
	//time.Sleep(time.Hour)

}

func taskTimer() {

	// 弃用标识符用法
	task.NewTaskChannel("init task")

	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("time = ", time.Now())
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

type NodeType int8

const (
	NodeTypeNull     NodeType = iota
	NodeTypeEdge              // 边缘设备
	NodeTypeGPUServe          // GPU服务器
)

// 边缘设备/GPU服务器节点信息
type (
	NodeInfo struct {
		NodeID        int          // 节点ID
		NodeType      NodeType     // 节点类型
		UpgradeRecord string       // 升级记录
		Os            string       // 操作系统
		Location      string       // 地理位置
		RegisterTime  string       // 注册时间
		Lists         []ServerInfo // 服务器列表信息
	}

	ServerInfo struct {
		CPUInfo           // cpu信息
		MemInfo           // 内存信息
		DiskInfo          // 磁盘信息
		LoadInfo int      // 负载情况
		IP       string   // IP地址
		Name     string   // 服务器名字
		Models   []string // 设备/服务器上跑的模型
	}

	DiskInfo struct {
		UsedPercent int // 利用率
		TotalMem    int // 总内存(GB)
		UsedMem     int // 已使用(GB)
	}

	MemInfo struct {
		UsedPercent int // 利用率
		TotalMem    int // 总内存(TB)
		UsedMem     int // 已使用(TB)
	}

	CPUInfo struct {
		UsedPercent int    // 利用率
		Framework   string // 架构
		Model       string // 型号
	}
)
