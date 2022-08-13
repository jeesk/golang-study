package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 32)
	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}

func searchByBaidu(str string) string {
	time.Sleep(1 * time.Second)
	return "baidu"
}
func searchByGoogle(str string) string {
	return "google"
}
func searchByBing(str string) string {
	return "bing"
}
