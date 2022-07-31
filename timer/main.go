package main

import (
	"fmt"
	"time"
)

func main() {
	// timer 和ticker 的区别是timer 是定时执行一次， 但是  ticker 是周期性执行。
	//newTickerTest()
	newTimerTest()
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
	timer := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case t := <-timer.C:
				fmt.Println("timerTest当前时间为:", t)
			}
		}
	}()

	for {
		time.Sleep(time.Second * 10)
	}
}
