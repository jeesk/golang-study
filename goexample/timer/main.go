package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer2 fird")
	}()
	stop := time2.Stop()
	if stop {
		fmt.Println("Timer 2 stopeed")
	}
	time.Sleep(2 * time.Second)

}
