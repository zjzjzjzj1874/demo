package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHamburgerWithCancel(t *testing.T) {
	HamburgerWithCancel()
}

func TestHamburgerWithTimeout(t *testing.T) {
	HamburgerWithTimeout()
}

func TestContextWithValue(t *testing.T) {
	ContextWithValue()
}

// context的超时控制机制nnnnnnnnnn
func TestContextWithValue2(t *testing.T) {
	ctx4, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("overslept")
	case <-ctx4.Done():
		fmt.Println("done :", ctx4.Err())
	}
}
