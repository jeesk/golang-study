package main

import (
	"fmt"
	"time"
)

//

func main() {
	testReadData()
}

func testTimer() {
	chan1 := make(chan int, 5)
	go func() {
		for i := 0; i < 100; i++ {
			chan1 <- 1
			fmt.Println("write data")
		}
	}()

	time.Sleep(100 * time.Second)
}

func testReadData() {
	chan1 := make(chan string, 2)
	for i := 0; i < 2; i++ {
		chan1 <- fmt.Sprint(i)
	}
	close(chan1)
	for i := 0; i < 3; i++ {
		fmt.Println(<-chan1)
	}
	fmt.Println("login")
	time.Sleep(5 * time.Second)
}
