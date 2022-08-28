package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then.Year())
	p(now)
	// 时间有sub 和add  api 用于增加或者时间
}
