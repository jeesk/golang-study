package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	singals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received messag ", msg)
	default:
		fmt.Println("no message receved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-singals:
		fmt.Println("received singal ", sig)
	default:
		fmt.Println("no activity")
	}

}
