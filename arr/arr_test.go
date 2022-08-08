package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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
func BenchmarkForTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100000; i++ {

	}
}
func BenchmarkForRangeTime(b *testing.B) {
	b.ResetTimer()
	var a [100000]int
	for range a {
	}

}

func Test_arr(t *testing.T) {
	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for i, v := range b {
		fmt.Println(i, v)
	}
	// 遍历的三种方式
	// range i
	for i := range a {
		fmt.Println(a[i])
	}
	// range iv
	for i, v := range a {
		fmt.Println(i, v)
	}
	// fori
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	// 忽略下标
	var times [2]int
	for range times {
		fmt.Println("迭代一次")
	}

	// 使用索引定义数组
	var s3 = [...]string{1: "世界", 0: "你好"}
	fmt.Println(s3)
	//数组可以定义: 字符串，结构体，函数,接口，管道

	_ = [2]string{"hello", "world"}

	var _ [1]image.Point
	var _ = [...]image.Point{image.Point{
		X: 0,
		Y: 0,
	}}
	// 定义函数
	var _ [2]func(reader io.Reader) (image.Image, error)
	var _ = [...]func(reader io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	var _ [2]interface{}
	var chanlist = [2]chan int{}
	fmt.Println(chanlist)

	// 使用空数组来强调操作 [0]int 表示空数组, c1 <- [0]int{} 表示放入一个空数组{} 表示空数组
	// struct{} 表示空结构器
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
		fmt.Println(c1)
	}()
	<-c1
	// [0]int，[3]int 和struct{} 是一种数据接口， 后面{} 表示结构器的值。
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2
	//
	fmt.Printf("b: %T ", s3)

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
