package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	//开启一个协程
	go func() {
		for {

			select {
			case num := <-ch:
				fmt.Println("get num is ", num)
			case <-time.After(2 * time.Second):
				fmt.Println("time's up!!!")
				//quit <- true
			}

		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(10 * time.Second)

	}

	time.Sleep(100 * time.Second)
}
