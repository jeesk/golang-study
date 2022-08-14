package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k"] = 7
	fmt.Println("map", m)
	v1 := m["k"]
	fmt.Println(v1)
	fmt.Println("map:", len(m))
	delete(m, "k")
	fmt.Println("map:", m)

	_, prs := m["k"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
