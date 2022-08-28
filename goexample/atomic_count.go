package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	// 使用atomic.AddInt 代替java 中胡说案子原子类型
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

}
