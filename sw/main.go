package main

import "fmt"

func main() {
	// 1. 匹配变量
	str := "hello"
	switch str {
	case "hello":
		fmt.Println("hello")
	default:
		fmt.Println("default")
	}

	// 匹配布尔表达式
	num := 10
	switch {
	case num > 10:
		fmt.Println("大于10")
	case num < 10:
		fmt.Println("小于10")
	case num == 10:
		fmt.Println("等于10")
	}

	type demo struct {
		name string
	}

	// 匹配struct
	d := demo{name: "123"}
	switch d {
	case demo{}:
		fmt.Println("demo")
	case demo{name: "123"}:
		fmt.Println("找到了")
	default:
		fmt.Println("default")
	}

}
