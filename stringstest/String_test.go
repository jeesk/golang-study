package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	s := "hello"
	fmt.Println(s[0:3])
	// 使用指针访问字符串的长度
	fmt.Println("s长度", len(s))
	fmt.Println("len(s2):",
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Len)

	// 字符串可以使用byte ,或者[]rune遍历
	// 字符串转byte, byte 转字符串， 主要使用 () 转换
	a := []byte("hello world")
	fmt.Println(a)
	fmt.Println(string(a))

}
