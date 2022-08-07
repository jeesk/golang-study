package main

import (
	"fmt"
	"testing"
)

const N = 10000

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "asdfasdf"
	}
	return s
}
func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}
func RangeFOrSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}
func RangeFOrRangeGet(s []string) {
	for i, _ := range s {
		_, _ = i, s[i]
	}
}
func BenchmarkForSlice(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRange(t *testing.B) {
	s := initSlice()
	t.ResetTimer()
	for i := 0; i >= t.N; i++ {
		RangeFOrSlice(s)
	}
}
func BenchmarkRangeGet(t *testing.B) {
	s := initSlice()
	t.ResetTimer()
	for i := 0; i >= t.N; i++ {
		RangeFOrRangeGet(s)
	}
}

// go test -bench. -run=None
func TestArr(t *testing.T) {
	fmt.Println("go")
	var a = [...]int{2, 2}
	for i := range a {
		fmt.Println(i)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 用数组range 迭代性能更好
}
