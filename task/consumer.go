package task

import (
	"fmt"
	"os"
	"time"
)

// goroutine with exit signal;and every 3 seconds print info;
func (tc *TaskChan) Consumer() {
	fmt.Printf("channel process pid:%d, ppid:%d .\n", os.Getpid(), os.Getppid())
	ok := true
	go func() {
		for ok {
			select {
			case v := <-tc.Task:
				fmt.Println("I am done:", v, " ; ", len(tc.Task))
			case <-time.After(time.Second * 3):
				fmt.Println("3s later:", time.Now().Format("2006-01-02 15:04:05"))
			case v := <-tc.ExitSignal:
				fmt.Println("exit signal:", v, ";3s later will exit channel")
				time.Sleep(3 * time.Second)
				ok = false
				fmt.Println("bye bye from pid:", os.Getpid())
			}
		}
	}()
}
