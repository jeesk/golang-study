package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["hello"] = "world"
	for s, value := range mp {
		fmt.Println(s, value)
	}

}
