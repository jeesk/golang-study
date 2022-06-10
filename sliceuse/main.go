package main

import (
	"fmt"
	"strconv"
	time2 "time"
)

// 数字转字符串
func main() {
	time := time2.Now()
	str1 := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		str1[i] = "hello" + strconv.Itoa(i)
		//fmt.Printf("str1[%d]=(%s)", i, str1[i])
	}
	fmt.Println("-----------cost times %d)", time2.Since(time))
	time = time2.Now()
	str2 := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		str2 = append(str2, "hello"+strconv.Itoa(i))
		//	fmt.Printf("str2[%d]=(%s)", i, str2[i])
	}
	//	fmt.Println(str2)
	fmt.Println("-----------cost times %d)", time2.Since(time))
}
