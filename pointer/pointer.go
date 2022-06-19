package main

import "fmt"

type Demo struct {
	name string
}

func main1() {
	demo := Demo{
		name: "songqifu",
	}
	demo1 := Demo{
		name: "songqifu1",
	}
	demo.Write(&demo1)
	fmt.Println(demo)
	fmt.Println(demo1)
}

func (p *Demo) Write(d *Demo) {
	slice := *p
	// Again as above.
	*p = *d
	*d = slice
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

type Person struct {
	name string
	age  int
}

func (r *Person) hello() string {
	return r.name
}

func main() {
	var r = &Person{"liu", 10}
	fmt.Printf("value=%v pointer= %p \n", r, r)
	fmt.Printf("value=%v pointer= %p \n", *r, *r)
	fmt.Printf("value=%v pointer= %p \n", *(*(&r)), &r)
}
