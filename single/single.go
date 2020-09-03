package single

import "sync"

type Single struct{}

var (
	once *sync.Once
	ins  *Single
)

// 单例模式实现
func GetIns() *Single {
	once.Do(func() {
		ins = &Single{}
	})

	return ins
}
