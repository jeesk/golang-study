package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v ", p)
	fmt.Println()
	fmt.Printf("struct2: %+v", p)
	fmt.Println()

	fmt.Printf("struct3 %#v", p)
	fmt.Println()

	fmt.Printf("type %T ", p)
	fmt.Println()
	fmt.Printf("bool: %t", true)
	fmt.Println()

	fmt.Printf("int: %d", 123)
	fmt.Println()
	fmt.Printf("%b", 12)

	fmt.Printf("%x ", 1)
	fmt.Println()

	fmt.Printf("%p ", p)
	fmt.Println()

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

}
