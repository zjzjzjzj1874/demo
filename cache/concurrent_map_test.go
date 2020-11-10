package cache

import (
	"fmt"
	"testing"
)

func TestConcurrentMapBenchmarkAdapter_Set(t *testing.T) {
	c := CreateConcurrentMapBenchmarkAdapter(16)
	c.Set("set", "hello world.")
	fmt.Println(c.Get("set"))
}
