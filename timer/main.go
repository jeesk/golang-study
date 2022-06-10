package main

import (
	"fmt"
	"time"
)

func main() {

	newTickerTest()
}

// newTimer 发送一次
func newTimerTest() {
	fmt.Println("timerTest1当前时间为:", time.Now())
	timer := time.NewTimer(1 * time.Second)
	go func() {

		for {
			t := <-timer.C
			fmt.Println("timerTest1当前时间为:", t)
		}
	}()
	for {
		time.Sleep(time.Second * 1)
	}

}

// 周期发送
func newTickerTest() {
	fmt.Println("timerTest当前时间为:", time.Now())
	timer := time.NewTicker(1 * time.Second)
	go func() {

		for {
			t := <-timer.C
			fmt.Println("timerTest当前时间为:", t)
		}
	}()
	for {
		time.Sleep(time.Second * 1)
	}
}
