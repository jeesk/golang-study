package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("di1", "name", "asdfas")
	fmt.Println("path is ", p)

	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("base(p)", filepath.Base(p))

	// https://gobyexample.com/directories
}
