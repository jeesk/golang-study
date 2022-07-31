package main

import (
	"fmt"
	"time"
)

func main() {
	// 计算时间
	doThing()
	getException()
}

// 捕捉异常
func getException() {
	defer func() {
		err := recover()
		if err != nil {
			// 使用runtime.Caller 获取异常的文件信息
			//caller, file, line, ok := runtime.Caller(0)
			fmt.Println(err)
		}
	}()
	if time.Now().Unix()%2 == 0 {
		panic("获取到了错误的东西")
	}
}

func doThing() {
	start := time.Now().Unix()
	defer func() {
		println(time.Now().Unix() - start)
	}()
	sum := 0
	for i := 1; i < 10000000000; i++ {
		sum /= i
	}
}
