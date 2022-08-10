package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {

	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)
	// 这种不能保证稳定的输出, 使用 ctrl + c 来实现
	time.Sleep(5 * time.Second)

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("quit (%v)", <-sig)
}
