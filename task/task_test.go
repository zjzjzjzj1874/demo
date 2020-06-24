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
