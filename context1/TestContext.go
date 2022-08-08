package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// 测试content WithValue 携带参数
	//testLogger()
	//testWithCancel()

	// 优雅宕机
	gracefulShutdown()

}

func testLogger() {
	// 使用content withValue 携带参数
	l := logger{}
	// 请求进来
	ctx := l.setReqInfo(context.Background(), "adsfasdfasdf")
	// 处理业务报错
	l.log(ctx, errors.New("{{requetId}} 服务调用发生熔断 "))
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

// 使用WithCancel 取消执行的操作。
func testWithCancel() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	msgC := make(chan string, 1000)
	for i := 0; i < 4; i++ {
		go doSomeThings(ctx, msgC)
	}
	//   dosomethins
	rand.Seed(100)
	randomNumber := rand.Intn(1000)
	for i := 0; i < 1000; i++ {

		if i == randomNumber {
			// 取消消费消息
			// 这里的取消消费也可以放在WithCancel 的下一行，直接defer withCancel() ， 这样当主线程执行完成后，就会取消消费消息的任务
			cancelFunc()
			break
		} else {
			msgC <- fmt.Sprint(i)
		}
	}
	time.Sleep(20 * time.Second)

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

// 使用WithTimeout 来处理优雅宕机

func NewContextWithTimeout(td time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), td)
}

func gracefulShutdown() {
	ctx, cancel := NewContextWithTimeout(5 * time.Second)
	// 服务优雅关机， 无论的连接是否断开 3秒后都直接宕机
	defer cancel()
	err := shutdown(ctx)
	if err != nil {
		fmt.Println("优雅宕机失败... ")
	} else {
		fmt.Println("优雅宕机成功... ")
	}
	fmt.Println("server extings ")
}

func shutdown(ctx context.Context) error {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			// 这里可以使用timer.Reset 一直做其他的事, 延迟结束的时间。
			// 假设这里还需要做其他的操作，可以设置  // timer.Reset(10 * time.Second)
			return nil
		}
	}
	return nil
}
