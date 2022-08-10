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
	// sekect 随机选择case 的发送和接受可以被立刻操作。 请教方浩成或者其他人。 这里的select 关键字使用
	var start = time.Now()
	select {
	// 这里为什么要使用时间等待呢？ select 关键字的理解。 如果下面两个都会被阻塞。直到有一个可以继续可以执行下去。
	// 假如这里阻塞， 切换到time.After 的时候， 是不是可以理解 这个管道其实状态已经被保存下来了。
	after1 := time.After(p.timeout)
	case sub <- v:
		fmt.Println("执行一次发送", v)
	case asdf := <- *after1 :
		fmt.Println("等待一次 耗时", time.Since(start), asdf)
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
