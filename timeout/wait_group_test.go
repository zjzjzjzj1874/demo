package timeout

import "testing"

func TestTimeoutWaitGroup(t *testing.T) {
	WaitGroupForTimeout(5)
}
