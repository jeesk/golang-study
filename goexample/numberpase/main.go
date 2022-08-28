package main

import (
	"fmt"
	"strconv"
)

func main() {
	float, _ := strconv.ParseFloat("10.0", 64)
	fmt.Println("float ", float)

	i, _ := strconv.ParseInt("12", 0, 64)
	fmt.Println("number is ", i)

	// number to string
	itoa := strconv.Itoa(121)
	fmt.Println(itoa)
}
