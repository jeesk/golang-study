package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Calling main
func main() {
	testTime()
}

func testTime() {

}

func testLimitReader() {
	// Defining r using NewReader
	r := strings.NewReader("G")

	// Calling LimitReader method with its parameters
	res := io.LimitReader(r, 3)

	// Calling Copy method with its parameters
	op, err := io.Copy(os.Stdout, res)

	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	fmt.Printf("\nn:%v\n", op)
}
