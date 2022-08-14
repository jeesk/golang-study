package main

import (
	"fmt"
	"math"
)

const s string = "constants"

func main() {
	fmt.Println(s)

	const n = 50000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(32))

	fmt.Println(math.Sin(n))
}
