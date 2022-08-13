package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	// 用来收消息
	subscriber chan interface{}         // 订阅者为一个管道
	topic      func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex         // 读写锁
	buffer      int                  // 订阅队列的缓存大小
	timeout     time.Duration        // 发布超时时间
	subscribers map[subscriber]topic // 订阅者信息
}

// 构建一个发布者对象, 可以设置发布超时时间和缓存队列的长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topic),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.addSubscribeTopic(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) addSubscribeTopic(topic topic) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic(
	sub subscriber, topic topic, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	var start = time.Now()

	after := time.After(p.timeout)
	select {
	//var after1 <-chan time.Time = time.After(p.timeout)

	case sub <- v:
		fmt.Println("执行一次发送", v)
	case v, ok := <-after:
		fmt.Println("等待一次 耗时", time.Since(start), v, ok)
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func block1() <-chan time.Time {
	fmt.Println("execute ")
	return time.After(10 * time.Second)
}
