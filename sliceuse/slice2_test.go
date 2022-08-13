package main

import (
	"testing"
)

func handle(data []int) {
	return
}

const N = 100

func getSlice() []int {
	a := []int{}
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			a = append(a, 0)
		} else {
			a = append(a, 1)
		}
	}
	return a
}

func BenchmarkDeleteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := DeleteSlice(getSlice())
		handle(data)
	}
}

func BenchmarkDeleteSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := DeleteSlice1(getSlice())
		handle(data)
	}
}

func BenchmarkDeleteSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := DeleteSlice2(getSlice())
		handle(data)
	}
}

func DeleteSlice(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

func DeleteSlice1(a []int) []int {
	ret := make([]int, 0, len(a))
	for _, val := range a {
		if val == 1 {
			ret = append(ret, val)
		}
	}
	return ret
}

func DeleteSlice2(a []int) []int {
	j := 0
	for _, val := range a {
		if val == 1 {
			a[j] = val
			j++
		}
	}
	return a[:j]
}
