// 求两个sync.map的差集
package main

import "sync"

func DifferenceMap(totalMap, subMap *sync.Map) (differenceKeys []interface{}) {
	totalMap.Range(func(key, value interface{}) bool {
		if _, ok := subMap.Load(key); !ok {
			differenceKeys = append(differenceKeys, key)
		}
		return true
	})
	return
}
