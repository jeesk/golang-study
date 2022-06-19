package main

import (
	"fmt"
	"strings"
)

const (
	i    = iota
	j    = iota
	k, l = iota, iota
)

func main() {
	fmt.Println(i, j, k, l)
	map1 := make(map[string]string)
	map1["hello"] = "1"
	map1["word"] = "2"
	for s, s2 := range map1 {
		fmt.Println(s, s2)
	}

}

func addSlice() {
	// 初始化为0
	myList := make([]int, 0)
	myList = append(myList, 1, 2, 3)
	fmt.Println(fmt.Sprint(myList))
	// 初始化为3, 手动赋值
	myList2 := make([]int, 3)
	myList2[0] = 0
	myList2[1] = 1
	myList2[2] = 2
	fmt.Println(fmt.Sprint(myList2))
	// 初始化为2， 使用append, 这个时候第一个值不会被覆盖
	myList3 := make([]int, 2)
	myList3 = append(myList3, 1, 2)
	fmt.Println(fmt.Sprint(myList3))
}

func arr2Slice() {
	arr := [5]string{"1", "2", "3", "4", "5"}
	println("打印指定索引", arr[1])
	println("打印全部", strings.Join(arr[:], ","))
	println("打印全部0~2，左开右闭", strings.Join(arr[0:2], ","))
	println("打印2~4，左开右闭", strings.Join(arr[2:4], ","))
	// 如果前面或者后面的索引是第一个或者最后一个，可以省略
	println("打印0~4，左开右闭", strings.Join(arr[:4], ","))
	println("打印2~5，左开右闭", strings.Join(arr[2:], ","))

}

func testArr() {
	// 第一种 申明数组，简介赋值
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	// 第二种 申明数组直接赋值， 这里的[2] 可以使用 ... 代替
	var arr2 [2]int = [2]int{1, 2}
	// var arr2 := [2]int{1, 2}， 也可以使用系统推导
	fmt.Println(arr2)
	// 第三种初始化数组{索引: 值}
	ints := [3]int{1: 9}
	fmt.Println(ints)
	// 直接初始化的时候可以使用 ... 代替个数， 下面表示最大有10个，第10个元素的值是100
	arr5 := [...]int{10: 100}
	fmt.Println(arr5)

}

func testSlice() {
	// 创建2个切片
	slice1 := []int{2, 2}
	slice2 := []int{2, 2}
	fmt.Println(slice1)
	slice1 = append(slice1, 1, 2, 3)
	// 切片加上... 变成可变数组
	ints := append(slice1, slice2...)
	fmt.Println(ints)
}
