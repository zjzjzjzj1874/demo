package _interface

import (
	"math/rand"
	"time"
)

type Decoder interface {
	Open(uri string) error
	Close() error
}

func NewDecode() Decoder {
	// 随机种子
	rand.Seed(time.Now().Unix())

	switch rand.Intn(100) % 2 {
	case 0:
		return &GPUDecode{}
	default:
		return &CPUDecode{}

	}
}
