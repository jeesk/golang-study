package main

import "fmt"

func main() {
	var a [2]int
	fmt.Println(a)
	a[1] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[1])

	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3}
	fmt.Println("dcl", b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
