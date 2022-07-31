package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("demo", "abc", "atest")
	flag.Parse()
	fmt.Println(*s)

	mp := make(map[string]int)
	mp["hello"] = 12
	// 遍历第一个是key, 第二个是value ，value 可以不写入
	for s := range mp {
		fmt.Println(s, s)
	}
	// 遍历slice 第一个是index,  第二个是value  , 可以不便利value
	slice2 := make([]map[string]int, 2, 2)
	for index := range slice2 {
		fmt.Println(index, index)
	}
	ints := make([]int, 2)
	ints[0] = 1
	ints[1] = 2
	fmt.Println(fmt.Sprint(ints))

	intsP := make([]*int, 0)
	intsP1 := make([]*int, 0)

	for index, value := range ints {
		fmt.Println("value: ", &value)
		intsP = append(intsP, &value)
		intsP1 = append(intsP1, &index)
	}
	// 发现value 的值是同一个， 这是因为golang 的range 遍历的变量的指针是一样的地址 。 所以不能对该值取指针操作
	fmt.Println("新的value")
	for i := range intsP {
		fmt.Println(*intsP[i])
	}
	// 修改inntsP 的内容
}
