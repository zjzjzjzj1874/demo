package task

func (tc *TaskChan) Producer() {
	tc.Task <- tc.Input + 1
}
