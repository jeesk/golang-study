package main

import (
	"fmt"
	"sync"
)

func main() {
	testWaitGroup()
}
func testWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go doSomeThing(1, &wg)
	go doSomeThing(2, &wg)
	wg.Wait()
}

func doSomeThing(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := 0
	for i := 0; i < 100000; i++ {
		b += i
	}
	fmt.Println("end")
}
