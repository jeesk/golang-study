package main

import (
	"fmt"
	"reflect"
)

func PrintType(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println("i am int")
		fmt.Println()
	case string:
		fmt.Println("i am string")

	}
}

func main() {
	// 隐式转换成interface{}
	PrintType(12)
	hello := "hello"
	switch interface{}(hello).(type) {
	case int:

		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
	TypeAssert()
}

func TypeAssert() {
	var hello interface{} = 12
	switch hello.(type) {
	case int:
		fmt.Println(reflect.TypeOf(hello))
		fmt.Printf("%T", hello)
	case string:
		fmt.Println(hello)
	}
}
