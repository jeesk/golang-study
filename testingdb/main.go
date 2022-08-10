package main

import (
	"fmt"
	"testing"
)

type Tb struct {
	testing.TB
}

func (p *Tb) Fatal(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case string:
			fmt.Println("我是一个字符串")
		}
	}
	fmt.Println("TB.Fatal disabled!")
}

func (p *Tb) Fatal1(args ...any) {
	for _, v := range args {
		switch v.(type) {
		case string:
			fmt.Println("我是一个字符串")
		}
	}
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = &Tb{}
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
