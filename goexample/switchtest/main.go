package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("wirte", i, "as")
	switch i {
	case 1:
		fmt.Println("one ")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three ")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend ")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after nonn")
	}
	whatAmL := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool ,type is %T", t)
			fmt.Println()
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't konw type")
		}
	}
	whatAmL(true)
	whatAmL(1)
	whatAmL("hey")
}
