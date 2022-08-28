package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"c", "b", "b"}
	sort.Strings(str)
	fmt.Println("Strings", str)

	ints := []int{6, 2, 3}
	sort.Ints(ints)
	fmt.Println("Ints", ints)
	// 是否已经排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("soted", s)
	fmt.Println("Sorted", s)

}
