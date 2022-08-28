package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestFunc(t *testing.T) {
	testMutex()
}

func testMutex() {

	stu := &Stu{
		name: "song",
		age:  "12",
		s:    Stu1{name: "stu1"},
	}
	stu.st1 = &Stu{

		name: "i amstu.st1",
		age:  "",
		s:    Stu1{},
		st1:  nil,
	}
	// 获取第一个
	fmt.Println((*Stu)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)))).name)
	i := *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + unsafe.Sizeof("")))
	// 获取第二个
	fmt.Println(i)
	// 获取第三个
	ci := (*Stu1)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + unsafe.Sizeof("") + unsafe.Sizeof("")))
	fmt.Println(ci.name)
	// 获取第四个
	ci1 := *(*Stu)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + unsafe.Sizeof("") + unsafe.Sizeof("") + unsafe.Sizeof(reflect.ValueOf(stu.s))))
	fmt.Println(ci1)
	// https://v2ex.com/t/871263#;  TODO 暂时先放过
	fmt.Println()
}

type Stu struct {
	name string
	age  string
	s    Stu1
	st1  *Stu
}

type Stu1 struct {
	name string
}
