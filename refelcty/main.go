package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {

	ExampleKind1()
}

type Stu struct {
	name string
	age  uint32
}

func ExampleKind() {

	// 通过valueOf(v)>kind 获取数据类型
	for _, v := range []any{&Stu{name: "demo", age: 12}, "hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}
}

// Transport constructs go-stream-muxer compatible connections.
type Transport interface {

	// NewConn constructs a new connection
	NewConn(c net.Conn, isServer bool) string
}

func ExampleKind1() {

	fmt.Println(reflect.ValueOf(nil).IsNil())

	// 通过valueOf(v)>kind 获取数据类型
	for _, v := range []any{Stu{name: "demoP", age: 12}, Stu{name: "demoNOP", age: 123}, "hi", 42, func(name string) bool {
		return true
	}} {
		// kind 是返回特种类型的类别
		fmt.Printf("reflect.ValueOf(v) ->%v  \n", reflect.ValueOf(v))

		fmt.Println("vlaue interface: -> %v", reflect.ValueOf(v).Interface())
		fmt.Printf("reflect.ValueOf(v).Type() ->%v \n", reflect.ValueOf(v).Type())
		fmt.Printf("reflect.ValueOf(v).Kind() ->%d \n", reflect.ValueOf(v).Kind())
		fmt.Printf("reflect.ValueOf(v).Type().Kind() ->%d \n", reflect.ValueOf(v).Type().Kind())
		fmt.Printf("reflect.TypeOf(v) ->%v \n", reflect.TypeOf(v))
		fmt.Printf("reflect.TypeOf(v).Kind() %d \n", reflect.TypeOf(v).Kind())

		fmt.Println("_________________________________________________________ \n")
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			/*fmt.Printf(v.String())*/
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			/*fmt.Printf(v)*/
		default:
			/*fmt.Printf("unhandled kind %s \n", v.Kind())*/
		}
	}

}
