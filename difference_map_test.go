// 求两个sync.map的差集
package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestDifferenceMap(t *testing.T) {
	var (
		totalMap sync.Map
		subMap   sync.Map
		notExist []interface{}
	)
	totalMap.Store("key1", "key1")
	totalMap.Store("key2", "key2")
	totalMap.Store("key3", "key3")
	//subMap.Store("key3", "key3")

	notExist = DifferenceMap(&totalMap, &subMap)

	fmt.Println("not exist :", notExist)

}
