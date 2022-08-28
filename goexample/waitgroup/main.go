package main

import (
	"fmt"
	"sync"
)

func worker(n int) {
	fmt.Println("%v is start", n)

	fmt.Println("%v is end", n)
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
