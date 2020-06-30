// This package is aim to show how to use channel;
// to achieve producer and consumer.
package task

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var taskChan *TaskChan

func init() {
	taskChan = NewTaskChan("init task")
}

func TestNewTaskChan(t *testing.T) {
	fmt.Printf("main process pid:%d, ppid:%d .\n", os.Getpid(), os.Getppid())
	taskChan.Consumer()

	for i := 0; i < 10; i++ {
		taskChan.Input = i
		taskChan.Producer()
		time.Sleep(time.Second)
	}

	go func() {
		select {
		case v := <-taskChan.ExitSignal:
			fmt.Println("exit signal:", v)
		case <-time.After(15 * time.Second):
			taskChan.ExitSignal <- 1
		}
	}()

	time.Sleep(time.Hour)
}

// for range读取channel
func TestNewTaskChan1(t *testing.T) {
	retry := func(i int) bool {
		ok := i%2 == 0
		return ok
	}
	go func() {
		// 场景：当需要不断从channel读取数据时
		// 原理：使用for-range读取channel，这样既安全又便利，当channel关闭时，for循环会自动退出，
		// 		无需主动监测channel是否关闭，可以防止读取已经关闭的channel，造成读到数据为通道所存储的数据类型的零值。
		for x := range taskChan.Task {
			fmt.Println(x)
			if retry(x) { // 添加过滤条件 ==> 满足条件可以忽略,不满足可以继续retry
				taskChan.Task <- x + 1
			}
		}
	}()

	for i := 0; i < 10; i++ {
		taskChan.Task <- i
		time.Sleep(time.Second * 3)
	}

	time.Sleep(time.Hour)
}
