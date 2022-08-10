package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 看不懂看不懂
	p := NewPublisher(100*time.Millisecond, 0)
	defer p.Close()

	// 增加一个订阅者， 订阅所有的信息
	all := p.Subscribe()

	// 订阅包含golang 的信息
	golang := p.addSubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
