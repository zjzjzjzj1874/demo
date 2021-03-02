package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 随机生成0-1的小数(如果没有上面的init方法,就是伪随机) ==> golang伪随机和真随机的区别(即种子数的区别)
func TestMyTestRand01(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(float64(rand.Intn(100)) / 100)
	}
}

// 切片测试(开闭区间测试)
func TestMyTestSlice(t *testing.T) {
	arr := make([]int, 10)

	fmt.Println(arr[0:4])
	fmt.Println(arr[0:5])
}

// IOT gateway测试
func TestIotGateway(t *testing.T) {
	// function: 生成IOT gateway需要的鉴权token
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authBytes := md5.Sum([]byte("HXBYcA32PgQhiHAsQB" + timestamp))
	fmt.Println(fmt.Sprintf("%s;%x;%s", "rockontrol", authBytes, timestamp))
}

// region waitGroup测试性能
// 不用waitGroup测试
func TestWithoutWaitGroup(t *testing.T) {
	intArr := make([]int, 20)
	var param string
	for idx := range intArr {
		param += fmt.Sprintf("%d,", idx)
	}
	postHttp(param)
}

// waitGroup用法测试
func TestWaitGroup(t *testing.T) {
	intArr := make([]int, 20)
	wg := sync.WaitGroup{}

	var (
		param string
		go0   string
		go1   string
		go2   string
		go3   string
	)
	for idx := range intArr {
		switch idx % 4 {
		case 0:
			go0 += fmt.Sprintf("%d,", idx)
		case 1:
			go1 += fmt.Sprintf("%d,", idx)
		case 2:
			go2 += fmt.Sprintf("%d,", idx)
		default:
			go3 += fmt.Sprintf("%d,", idx)
		}
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		switch i {
		case 0:
			param = go0
		case 1:
			param = go1
		case 2:
			param = go2
		case 3:
			param = go3
		}
		go postHttpWithWaitGroup(param, &wg)
	}

	wg.Wait()
	fmt.Println("all work has done!")
}

func postHttpWithWaitGroup(param string, group *sync.WaitGroup) {
	defer group.Done()
	resp, err := http.Get("http://localhost:8080/hello?param=" + param)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	fmt.Println(string(content))
}

func postHttp(param string) {
	resp, err := http.Get("http://localhost:8080/hello?param=" + param)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http get err [err:%s]\n", err.Error())
		return
	}
	fmt.Println("postHttp : ", string(content))
}

// endregion waitGroup测试性能

// region 函数耗时测试
func TestTimeUsedCount(t *testing.T) {
	begin := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
	used := time.Since(begin)
	now := time.Now()
	sub := now.Sub(begin)
	fmt.Println("use time ==> ", used, " ; sub time ==> ", sub)

	// 优雅的写法
	defer timeCost()() //注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
}

//@brief：耗时统计函数
func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}

// endregion 函数耗时测试

// 浮点数与字符串转化
func TestStringParseToFloat64(t *testing.T) {
	tagVal, err := strconv.ParseUint("9765439183791093097", 10, 64)
	fmt.Println(tagVal, err)
}

// 是否是2的幂
func TestPowerOfTwo(t *testing.T) {
	var x uintptr
	x = 2
	fmt.Println(x&(x-1) == 0)
	x = 4
	fmt.Println(x&(x-1) == 0)
	x = 6
	fmt.Println(x&(x-1) == 0)
	x = 8
	fmt.Println(x&(x-1) == 0)
}

// os包中的pid和getPid
func TestPidInOs(t *testing.T) {
	fmt.Println("gid : ", syscall.Getegid())
	fmt.Println("uid : ", syscall.Geteuid())
	fmt.Println("uid : ", syscall.Getgid())
	fmt.Println("ppid : ", syscall.Getppid())
	fmt.Println("pid : ", syscall.Getpid())

	pgid, err := syscall.Getpgid(syscall.Getpid())
	fmt.Println("pgid : ", pgid, err)
	go func() {
		fmt.Println("go1 gid : ", syscall.Getegid())
		fmt.Println("go1 uid : ", syscall.Geteuid())
		fmt.Println("go1 uid : ", syscall.Getgid())
		fmt.Println("go1 ppid : ", syscall.Getppid())
		fmt.Println("go1 pid : ", syscall.Getpid())

		pgid, err := syscall.Getpgid(syscall.Getpid())
		fmt.Println("go1 pgid : ", pgid, err)
		time.Sleep(time.Hour)
	}()
	time.Sleep(time.Hour)
	// todo  开启协程:用waitGroup的来控制多个时间 ==> ai-agent中请求机器的方法拆分开来 避免超时
}

func TestTypeAndBit(t *testing.T) {
	t.Run("数值类型占用字节空间", func(t *testing.T) {
		var i1 int = 1   // int是最特殊的,和当前程序运行的平台强相关 ==> 本机是Mac,占用8个字节
		var i2 int8 = 2  // 8位占用一个字节
		var i3 int16 = 3 // 16位占用2个字节
		var i4 int32 = 4 // 32位占用4个字节
		var i5 int64 = 5 // 64位占用8个字节
		fmt.Println(unsafe.Sizeof(i1))
		fmt.Println(unsafe.Sizeof(i2))
		fmt.Println(unsafe.Sizeof(i3))
		fmt.Println(unsafe.Sizeof(i4))
		fmt.Println(unsafe.Sizeof(i5))
	})

	t.Run("生成随机字符串（包含大小写字母+数字）", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println(CharStringByASCII(6))
		}
	})

	t.Run("生成数字型随机字符串", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second)
			fmt.Println(NumString(6))
		}
	})
	time.Sleep(time.Hour)
}

func BenchmarkBytesToStr(b *testing.B) {
	b.Run("ASCII码生成", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CharStringByASCII(6)
		}
	})
	b.Run("随机字符串生成", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RandStringRunes(6)
		}
	})
}

func CharStringByASCII(len int) string {
	/*生成随机字符串（包含大小写字母+数字）*/
	kinds, result := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, len)
	for i := 0; i < len; i++ {
		ikind := rand.Intn(3)
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NumString(len int) string {
	/*生成数字型随机字符串*/
	r := rand.New(rand.NewSource(time.Now().Unix()))
	str := ""
	for i := 0; i < len; i++ {
		b := r.Intn(10)
		if i == 0 && b == 0 {
			b = 1
		}
		str += strconv.Itoa(b)
	}
	return str
}
