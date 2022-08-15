package main

import "time"

func main() {
	ticker := time.NewTicker(500 * time.Microsecond)
	done := make(chan bool)
	go func() {
		// https://gobyexample.com/tickers
	}()
}
