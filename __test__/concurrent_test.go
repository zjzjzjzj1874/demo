package __test__

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTaskMap(t *testing.T) {
	for i := 0; i < 200; i++ {
		go func(a int) {
			taskMap.Store(a, a)
			//GetTaskMap().Store(a, a)
		}(i)
	}

	time.Sleep(time.Second)
	i := 0
	GetTaskMap().Range(func(key, value interface{}) bool {
		i++
		fmt.Printf("key:%v,value:%v \n", key, value)
		return true
	})
	fmt.Println(i)

}
