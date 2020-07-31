package main

import (
	"demo/task"
	"fmt"
	"time"
)

func main() {
	// todo  开启协程:用waitGroup的来控制多个时间 ==> ai-agent中请求机器的方法拆分开来 避免超时



	//tagVal, err := strconv.ParseUint("9765439183791093097", 10, 64)
	//fmt.Println(tagVal, err)

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

	//m := make(map[string]CPUInfo)
	//m["1"] = CPUInfo{
	//	UsedPercent: 50,
	//	Framework:   "arm",
	//	Model:       "i7",
	//}
	//m["2"] = CPUInfo{
	//	UsedPercent: 60,
	//	Framework:   "arm",
	//	Model:       "i5",
	//}
	//
	//for k, v := range m {
	//	v.UsedPercent += 10
	//	m[k] = v
	//}
	//time.Sleep(time.Second)
	//fmt.Println(m)
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
