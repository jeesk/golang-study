package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var total uint64

// 使用原子类执行加减法
func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func Test_aDD(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}
