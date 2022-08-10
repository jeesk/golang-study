package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
}

func (p Person) Say() {
	fmt.Println(" i am person ", p.name)
}

type Stu struct {
	Person
	name string
	Flyer
}

func (f Stu) fly(str string) string {
	return "我会飞" + str
}

type Flyer interface {
	fly(str string) string
}

func GetValue(p Person) {
	fmt.Println(p.name)
}
func GetValue2(s Stu) {
	fmt.Println(s.name)
}

func (s Stu) say() {

	fmt.Println(" i am student ", s.name)
}

func main() {
	var stu = Stu{name: "xiaoming", Person: Person{
		name: "person",
	}}
	stu.Say()
	stu.say()

	var tb Flyer = Stu{
		Person: Person{},
		name:   "",
		Flyer:  nil,
	}
	tb.fly("asdfasdfsa")

	// 无法调用， 说明golang 的继承只表现在调用，作为参数的时候无法体现, 只有直接调用匿名字段后，再作为参数
	//GetValue(stu)
	GetValue(stu.Person)
	GetValue2(stu)

	people := make([]Person, 0)
	people = append(people, Person{name: "1"}, Person{name: "2"}, Person{name: "3"})
	Serve(people)
	time.Sleep(time.Second * 20)
}

func Serve(queue []Person) {
	for _, value := range queue {

		value := value
		go func() {
			process(&value) // 这儿有 Bug，解释见下。
		}()
	}
}
func process(person *Person) {
	fmt.Printf("%p ", person)
}
