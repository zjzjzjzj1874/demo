package task

type TaskChan struct {
	Desc       string
	Input      int
	Task       chan int
	Output     string
	ExitSignal chan int
}

func NewTaskChan(desc string) *TaskChan {
	tc := &TaskChan{Desc: desc}
	tc.Task = make(chan int, 100)
	tc.ExitSignal = make(chan int, 1)
	return tc
}

// deprecated
func NewTaskChannel(desc string) *TaskChan {
	tc := &TaskChan{Desc: desc}
	tc.Task = make(chan int, 100)
	tc.ExitSignal = make(chan int, 1)
	return tc
}
