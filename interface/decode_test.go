package _interface

import (
	"fmt"
	"testing"
)

func TestNewDecode(t *testing.T) {
	decode := NewDecode()
	switch decode.(type) {
	case *CPUDecode:
		fmt.Println("cpu decode")
	case *GPUDecode:
		fmt.Println("gpu decode")
	}
}
