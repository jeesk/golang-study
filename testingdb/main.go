package main

import (
	"fmt"
	"testing"
)

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")

	var i = IF(true, 1, 2)
	fmt.Println(i)

}

func IF[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
