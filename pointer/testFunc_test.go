package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestFunc(t *testing.T) {
	testMutex()
}

// 测试指针操作
func testMutex() {
	var l sync.Mutex
	p := unsafe.Pointer(&l)
	up0 := uintptr(p)
	// 获取第一个参数
	ps := (*int32)(p)
	*ps = 1

	p = unsafe.Pointer(up0 + uintptr(4))

	pm := (*int32)(p)
	*pm = 100
	l.Unlock()
	fmt.Println("unlocked")

	stu := &Stu{
		name: "song",
		age:  12,
	}

	// 获取第一个
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)))))
	i := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + unsafe.Sizeof("")))
	// 获取第二个
	fmt.Println(i)

}

type Stu struct {
	name string
	age  int
}
