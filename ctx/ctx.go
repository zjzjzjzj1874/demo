package ctx

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 场景1:一个人每秒吃0-5个汉堡,请问吃10个汉堡花几秒钟?
func HamburgerWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := eatHamburgerWithCancel(ctx)
	for n := range eatNum {
		if n >= 10 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果。。。")
	time.Sleep(1 * time.Second)
}

func eatHamburgerWithCancel(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
			}
		}
	}()

	return c
}

// 场景2:一个人吃汉堡,每秒钟吃0-5个汉堡,请问10s他可以吃多少个汉堡?
func HamburgerWithTimeout() {
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	chiHanBao(ctx)
	defer cancel()
}

func chiHanBao(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			incr := rand.Intn(6)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}

// 在ctx中携带对应的值,传递下去
func ContextWithValue() {
	ctx1 := context.WithValue(context.Background(), "greet", "hello world")
	//ctx2 := context.WithValue(ctx1, "session", 123456)
	//ctx3 := context.WithValue(ctx2, "session", 123)
	//fmt.Println(ctx3.Value("session")) // ctx.value在设计初就是不可变模式=>即设定完毕之后,ctx的所有信息都不可变了.

	//process(ctx)
	//
	//ses := ctx.Value("session")
	//fmt.Println(ses)
	//fmt.Println(reflect.TypeOf(ses))
	//ses = reflect.ValueOf(ses)
	//fmt.Println(ses)
	//fmt.Println(reflect.TypeOf(ses))

	// context配合select做超时处理  ==> 最常用于RPC超时调用
	ctx4, cancel := context.WithTimeout(ctx1, 1*time.Second)
	defer cancel()
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("overslept")
	case <-ctx4.Done():
		fmt.Println(ctx4.Err())
	}
}

func process(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	fmt.Println(ok)
	if !ok {
		fmt.Println("something wrong")
		return
	}

	if session != 123456 {
		fmt.Println("session 未通过")
		return
	}

	traceID := ctx.Value("trace_id")
	fmt.Println("traceID:", traceID, "-session:", session)
}
