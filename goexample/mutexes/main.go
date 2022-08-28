package main

import (
	"fmt"
	"sync"
)

type Container struct {
	my        sync.Mutex
	cotunters map[string]int
}

func (c *Container) inc(name string) {
	c.my.Lock()

	defer c.my.Unlock()
	c.cotunters[name]++
}
func main() {
	c := Container{
		cotunters: make(map[string]int),
	}

	var wg sync.WaitGroup
	doinc := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doinc("a", 1000)
	go doinc("a", 1000)
	go doinc("c", 1000)
	wg.Wait()
	fmt.Println(c.cotunters)
}
