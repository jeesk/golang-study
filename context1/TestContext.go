package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

func main() {
	l := logger{}
	// 请求进来
	ctx := l.setReqInfo(context.Background(), "adsfasdfasdf")
	// 处理业务报错
	l.log(ctx, errors.New("{{requestId}} 服务调用发生熔断 "))

}

// 使用ctx 传递参数
const contextLogKey = ("miniolog")

type logger struct{}

func (log logger) setReqInfo(ctx context.Context, requestInfo string) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, contextLogKey, requestInfo)
}

func doSomeThings(ctx context.Context, c chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("被取消消费动作")
			return
		case msg := <-c:
			fmt.Println("收到消息处理", msg)
		}
	}
}

func (log logger) log(ctx context.Context, err error) {
	if err == nil {
		return
	}
	value := ctx.Value(contextLogKey)
	if len(value.(string)) > 0 {
		if err != nil {
			s := err.Error()
			// 替换占位符号（这里简单的写一下，替换占位符号）
			fmt.Println(strings.ReplaceAll(s, "{{requestId}}", value.(string)))
		}
	}
}
