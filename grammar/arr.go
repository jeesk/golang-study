package main

import (
	"crypto/subtle"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func printType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Print("param is int")
	case string:
		fmt.Print("param is string")
	}
}

func main() {
	fmt.Print(subtle.ConstantTimeCompare([]byte("User"), []byte("User")))

	a := 10

	printType(a)
	// this is first way
	switch interface{}(a).(type) {
	case int:
		fmt.Print("param is int")
	case string:
		fmt.Print("param is string")
	}

	// this is second way
	var b = 10
	c := interface{}(b)
	switch c.(type) {
	case int:
		fmt.Print("param is int")
	case string:
		fmt.Print("param is string")
	}
	// this is three way

	var f interface{} = 23

	switch e := f.(type) {
	case int:
		fmt.Printf("param is int %v \n\r", e)
	case string:
		fmt.Printf("param is int %v \n\n", e)
	}

	person := Person{Name: "song", Age: 12, Addr: "wuhan"}
	fmt.Printf("person is can weite %v", reflect.ValueOf(person).CanSet())
	fmt.Printf("person is can weite %v", reflect.ValueOf(&person).Elem().CanSet())
	// newPerson := Person{Name: "song", Age: 12, Addr: "wuhan"}

	var numList []int = []int{1, 2}

	v1 := reflect.ValueOf(numList)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)

	// Slice 函数接收两个参数
	v2 := v1.Slice(0, 2)
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
}

func myarr() {
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	routes := [...]string{
		"/hi",
		"/contact",
		"/co",
		"/c",
		"/a",
		"/ab",
		"/doc/",
		"/doc/go_faq.html",
		"/doc/go1.html",
		"/α",
		"/β",
	}
	fmt.Print(arr1)
	fmt.Print(arr2)
	fmt.Print(routes)
}
func slice() {
	var slice1 []int
	slice2 := []int{1, 2}
	fmt.Print(slice1)
	fmt.Print(slice2)

	slice1 = append(slice1, slice2...)
	// if appended datalen > cap, so
}

func test2SliceWay() {
	picture := make([][]uint8, 5) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.

	for i := range picture {
		picture[i] = make([]uint8, 5)
	}
}

func test2sliceWay2() {
	picture := make([][]uint8, 5)
	pixels := make([]uint8, 5*5)
	for i := range picture {
		// every picels remove five eles
		picture[i], pixels = pixels[:5], pixels[5:]
	}
}

type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = slice

}
