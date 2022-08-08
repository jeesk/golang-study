package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// slice init way
	b := []int{}
	fmt.Println(b)
	c := []int{1, 2, 4}
	fmt.Println(c)
	// two ele , cap 3
	d := c[:2]
	fmt.Println(d)
	// two ele , cap is three
	e := c[0:2:3]
	fmt.Println(e)
	// cap is three, 0 cap, 0ele
	f := c[:0]
	fmt.Println(f)
	// TODO ------ create slice two way ,first direct 	assignment , two way use make
	// create null slice
	// 3 len, 3cap , 0ele
	g := make([]int, 0, 3)
	fmt.Println(g)
	// 3 ele, len and cap is 3
	h := make([]int, 3)
	fmt.Println(h)
	// cap is three , len is 2
	i := make([]int, 2, 3)
	fmt.Println(i)
	// range is same as array
	var slice1 []int
	slice1 = append(slice1, 2)
	// add slice need to unpackage
	slice1 = append(slice1, []int{1, 2, 4}...)

	// 不仅可以在尾部， 也可以 在头部 add
	slice1 = append([]int{1, 2}, slice1...)
	fmt.Println(slice1)

	a := []int{1, 2, 4}
	a = append(a, 0)
	fmt.Println(a)
	var x int32 = 1
	fmt.Println("a[x+1:]: ", a[x+1:])
	fmt.Println(a[x:])
	copy(a[x+1:], a[x:])

	fmt.Println(a)
	a[x] = 100
	fmt.Println(a)

}
