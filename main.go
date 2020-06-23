package main

import (
	task2 "demo/task"
	"fmt"
	"os"
	"time"
)

var task *task2.TaskChan

func init() {
	task = task2.NewTaskChan("init task")
}

func main() {
	fmt.Printf("main process pid:%d, ppid:%d .\n", os.Getpid(), os.Getppid())
	task.Consumer()

	for i := 0; i < 10; i++ {
		task.Input = i
		task.Producer()
		time.Sleep(time.Second)
	}

	go func() {
		select {
		case v := <-task.ExitSignal:
			fmt.Println("exit signal:", v)
		case <-time.After(15 * time.Second):
			task.ExitSignal <- 1
		}
	}()

	time.Sleep(time.Hour)
}
