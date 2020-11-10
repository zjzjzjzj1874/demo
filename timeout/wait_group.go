// waitGroup的超时处理机制
package timeout

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupForTimeout(sec int) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	go func(timeout int) {
		time.Sleep(time.Duration(timeout) * time.Second)
		wg.Done()
	}(sec)

	timeout := 10 * time.Second
	fmt.Printf("wait for wait group (up to %ds)\n", timeout)
	select {
	case <-done:
		fmt.Printf("wait group finished \n")
	case <-time.After(timeout):
		fmt.Printf("timeout waitting for wait group\n")
	}
	fmt.Printf("free at last\n")
}
