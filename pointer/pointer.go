package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

type Demo struct {
	name string
}

func main1() {
	demo := Demo{
		name: "songqifu",
	}
	demo1 := Demo{
		name: "songqifu1",
	}
	demo.Write(&demo1)
	fmt.Println(demo)
	fmt.Println(demo1)
}

func (p *Demo) Write(d *Demo) {
	slice := *p
	// Again as above.
	*p = *d
	*d = slice
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

type Person struct {
	name string
	age  int
}

func (r *Person) hello() string {
	return r.name
}

/*https://www.jianshu.com/p/7b3638b47845*/
func main() {

}

type TestPointer struct {
	b *Person // 私有变量
	A int
	m sync.RWMutex
	c *Person
}

func (T *TestPointer) OouPut() {
	fmt.Println("TestPointer OouPut:", T.A, T.b)

}

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

func string2bytetest() {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		_ = string2bytes("hello")
	}
	fmt.Println(time.Since(start))
}
func strin2bytelib() {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		_ = []byte("hello")
	}
	fmt.Println(time.Since(start))
}
