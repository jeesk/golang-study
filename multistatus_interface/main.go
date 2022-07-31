package main

import "fmt"

type Person interface {
	name() string
}
type Stu struct {
	Person
	Name string
}

func (s Stu) name() string {
	return s.Name
}
func GetValue(p Person) {
	fmt.Println(p.name())
}
func GetValue2(s Stu) {
	fmt.Println(s.Name)
}

func (s Stu) say() {
	fmt.Println(" i am student ", s.name)
}
func main() {
	var stu Person = Stu{Name: "songqifu"}
	GetValue2(stu.(Stu))
	GetValue(stu)

	var stu1 = Stu{Name: "songqifu1"}

	GetValue2(stu1)
	GetValue(stu1)
}
