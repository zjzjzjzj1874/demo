package __test__

import "sync"

var (
	taskMap = &sync.Map{}
)

func GetTaskMap() *sync.Map {
	return taskMap
}
