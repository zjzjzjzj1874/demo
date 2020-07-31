package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMyTestErr(t *testing.T) {
	fmt.Println(time.Unix(1595487025, 0).Format(time.RFC3339))

	years, month, day := time.Now().Date()
	fmt.Println(time.Date(years, month, day, 0, 0, 0, 0, time.FixedZone("CST", 8*60*60)).Format(time.RFC3339))
	fmt.Println(time.Date(years, month, day, 0, 0, 0, 0, time.UTC).Format(time.RFC3339))
}

// 切片在for range时设置值,是否会成功
func TestSliceForRange(t *testing.T) {
	m := make(map[string]CPUInfo)
	m["1"] = CPUInfo{
		UsedPercent: 50,
		Framework:   "arm",
		Model:       "i7",
	}
	m["2"] = CPUInfo{
		UsedPercent: 60,
		Framework:   "arm",
		Model:       "i5",
	}

	for k, v := range m {
		v.UsedPercent += 10
		m[k] = v
	}
	fmt.Println(m)
}
