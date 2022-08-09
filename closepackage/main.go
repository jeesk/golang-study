package main

import (
	"fmt"
)

func main() {
	fmt.Println(TestDefer())

	//
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

func TestDefer() (v int) {
	// golang对外部变量的捕捉是通过引用的方式来获取的。  闭包必须注意， 是否引用的变量。
	defer func() {
		v++
	}()
	return 42
}
