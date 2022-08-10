package main

import "utils"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
	utils.GetLength("asdfsadf")
}
