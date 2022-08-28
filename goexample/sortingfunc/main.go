package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 实现了 less, len,swap 的函数可以自定义排序
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
