package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01 22:08:41")
	p(t1)
}
