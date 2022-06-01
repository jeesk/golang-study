package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
}

func main() {
	person := Person{name: "song"}
	rValue := reflect.ValueOf(person)
	obj := reflect.New(rValue.Type()).Interface()
	fmt.Printf("%p", obj)
}
