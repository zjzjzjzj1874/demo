package cache

import (
	"fmt"
	"testing"
)

var emp = NewExpiredMap()

func TestNewExpiredMap(t *testing.T) {
	//emp.Set("key1", "key1", 3)
	//emp.Set("key2", "key2", 3)
	//emp.Set("key2", "key3", 3)
	//fmt.Println(emp.Get("key1"))
	//fmt.Println(emp.Get("key2"))
	//time.Sleep(time.Hour)
	//
	emp.DoForEach(func(i interface{}, i2 interface{}) {
		fmt.Println("in : ", i, i2)
	})
	emp.DoForEachWithBreak(func(i interface{}, i2 interface{}) bool {
		fmt.Println("DoForEachWithBreak : ", i, i2)
		return true
	})
}
