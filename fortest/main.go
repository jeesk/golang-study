package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("demo", "abc", "atest")
	flag.Parse()
	fmt.Println(*s)

	mp := make(map[string]string)
	mp["hello"] = "world"
	for s, value := range mp {
		fmt.Println(s, value)
	}

}
