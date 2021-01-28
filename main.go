package main

import (
	"demo/task"
	"fmt"
	"time"
)

var (
	a   bool = true
	CST      = time.FixedZone("CST", 8*60*60)
)

type MyTest struct {
	A   int `json:"a,default:100"`
	Key string
	B   int
	Bl  bool `json:"bl,default:true"`
}

func modifyMyTest(test MyTest) {
	//if 1 > test.B {
	//	test.B = 1
	//}
	//if 100 > test.A {
	//	test.A = 100
	//}
}

func main() {

	a := uint64(465)
	b := uint64(467)
	fmt.Println(a - b)
	fmt.Println(int64(a) - int64(b))
	fmt.Println(uint64(int64(a) - int64(b)))

	//test := MyTest{}

	//a := (*MyTest)(&test)
	//a.A = 100
	//fmt.Println(a)
	//_ = json.Unmarshal([]byte("{\"Key\":\"hello\",\"B\":0}"), &test)
	//fmt.Println(test)
	//modifyMyTest(test)
	//fmt.Println(test)
	//
	//啊 := 2.9808080808012
	//啊, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", 啊), 64)
	//fmt.Println(啊)

	//rand.Seed(time.Now().Unix())
	//ch := make(chan string)
	//go func() {
	//
	//	i := rand.Intn(5) % 2
	//	fmt.Println("i = ", i)
	//	if i == 0 {
	//		ch <- "ok1"
	//	}
	//	time.Sleep(1 * time.Second)
	//	ch <- "ok2"
	//
	//}()
	//select {
	//case <-time.After(time.Second):
	//	fmt.Println("超时")
	//case <-ch:
	//	//fmt.Println(res)
	//}
	//fmt.Println(time.Now().In(CST).Add(-60 * time.Minute).Hour())
	//fmt.Println(time.Now().In(CST).Hour())

	//start := time.Now()
	//fmt.Println(0 % 100)
	//time.Now().Second()
	//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	//fmt.Println(time.Now().Sub(start).Seconds())

	// 用这个来控制
	//m := sync.Map{}
	//queue := list.New()
	//m.Store(1, queue)
	//
	//for i := 0; i < 20; i++ {
	//	val, ok := m.Load(1)
	//	if ok {
	//		q := val.(*list.List)
	//		if q.Len() > 10 {
	//			ele := q.Back()
	//			if ele != nil {
	//				q.Remove(ele)
	//				q.PushFront(i)
	//			}
	//		} else {
	//			q.PushFront(i)
	//		}
	//		m.Store(1, q)
	//
	//		for iter := q.Back(); iter != nil; iter = iter.Prev() {
	//			fmt.Println("item:", iter.Value)
	//		}
	//	}
	//}
	//
	//// region 测试图片
	//resp, err := http.Get("https://static001.infoq.cn/resource/image/07/95/070aeb1b3ab22eaf69d5f59dfa622495.png")
	//if err != nil {
	//	fmt.Println("get picture err:", err)
	//	return
	//}
	//fileData, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("read picture err:", err)
	//	return
	//}
	//
	//// 存图片
	//out, _ := os.Create("/Users/jiahua-zhoujian/Tools/test/pic/test.jpeg")
	//_, _ = io.Copy(out, bytes.NewReader(fileData))
	// endregion 测试图片

	//m1 := new(map[string]int)
	//(*m1)["hello"] = 2
	//m1 := make(map[string]int)
	//m1["hello"] = 2
	//fmt.Println(m1)

	//num, err := strconv.ParseInt("", 10, 64)
	//if err != nil {
	//	fmt.Printf("err:%v \n", err)
	//}
	//fmt.Printf("num:%d \n", num)
	//
	//imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	//////图片正则
	////reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	////name := reg.FindStringSubmatch(imagPath)[0]
	////fmt.Print(name)
	////通过http请求获取图片的流文件
	//resp, _ := http.Get(imagPath)
	//body, _ := ioutil.ReadAll(resp.Body)
	//out, _ := os.Create("/Users/jiahua-zhoujian/Tools/test/pic/test.jpeg")
	//io.Copy(out, bytes.NewReader(body))
	//
	//// region context
	//// https://www.infoq.cn/article/fIBEaRLQstWIEkd94BCR
	//// Create an HTTP server that listens on port 8000
	//_ = http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	ctx := r.Context()
	//	// This prints to STDOUT to show that processing has started
	//	_, _ = fmt.Fprint(os.Stdout, "processing request\n")
	//	// We use `select` to execute a peice of code depending on which
	//	// channel receives a message first
	//	select {
	//	case <-time.After(2 * time.Second):
	//		// If we receive a message after 2 seconds
	//		// that means the request has been processed
	//		// We then write this as the response
	//		_, _ = w.Write([]byte("request processed"))
	//	case <-ctx.Done():
	//		// If the request gets cancelled, log it
	//		// to Stdout
	//		_, _ = fmt.Fprint(os.Stdout, "request cancelled\n")
	//	}
	//}))
	// endregion context

	// region 协程中,变量的微小变化
	//i := 0
	//for {
	//	if i%2 != 0 {
	//		i++
	//		continue
	//	}
	//	//c := i
	//	go func() {
	//		fmt.Println("int2 === ", i)
	//	}()
	//	i++
	//
	//	time.Sleep(time.Second)
	//}
	// endregion 协程中,变量的微小变化

	//time.Sleep(time.Hour)

	//params := "/lb-ai/v1/rt-analyze/uniform-exclusive"
	//secret := "f4dce4ad3b13eb26f6f3ceb29e5629a5fcf24221a28c640ba2a53605fc328357"
	//timeStamp := 1601350600
	//newString := params + secret + fmt.Sprintf("%d", timeStamp)
	//strByte := []byte(newString)
	//
	//// 方法1
	//hash := sha256.New()
	//_, _ = hash.Write(strByte)
	//sign := hex.EncodeToString(hash.Sum(nil))
	//fmt.Println(sign)
	//
	//// 方法2
	//h := sha256.New()
	//h.Write([]byte(newString))
	//sec := h.Sum(nil)
	//fmt.Printf("%x\n", sec)
	//
	//// 方法3
	//sum := sha256.Sum256([]byte(newString))
	//fmt.Printf("%x", sum)

	// 位运算解决签到方案(每个月最多31天,in32就够用了,然后签到一天,将当天的位设置为1即可)
	//const (
	//	a = 1 << iota // 1 << 0
	//	b             // 1 << 1
	//	c             // 1 << 2
	//	d             // 1 << 3
	//	e             // 1 << 4
	//	f             // 1 << 5
	//)
	//
	//signIn := 0
	//signIn |= a
	//signIn |= d
	//fmt.Println(a, d, b, c, d, e, f, 1<<1, 3<<2)
	//fmt.Println(" sign in:", signIn, signIn&a == 0, signIn&b == 0, signIn&c == 0, signIn&d == 0, signIn&e == 0, signIn&f == 0)

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
