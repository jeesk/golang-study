package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// Config 配置
type Config struct {
	Listen  uint16
	Forward []string
}

// DoServer 启动服务器
func DoServer(config []Config) {
	var handle = func(cfg Config) {
		// 负载均衡封装
		var getforward func() string
		var fid = -1
		if len(cfg.Forward) > 1 {
			getforward = func() string {
				fid++
				if fid >= len(cfg.Forward) {
					fid = 0
				}
				return cfg.Forward[fid]
			}
		} else {
			getforward = func() string {
				return cfg.Forward[0]
			}
		}
		var doconn = func(conn net.Conn) {
			// 处理进来的连接
			defer conn.Close()
			forw := getforward()
			log.Println(forw)
			fconn, err := net.Dial("tcp", forw)
			if err != nil {
				log.Println(err)
				return
			}
			defer fconn.Close()
			go io.Copy(conn, fconn)
			io.Copy(fconn, conn)
		}
		// 处理
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", cfg.Listen))
		if err != nil {
			panic(err)
		}
		defer lis.Close()
		log.Println("Listen on", cfg.Listen)
		for {
			conn, err := lis.Accept()
			if err != nil {
				continue
			}
			go doconn(conn)
		}
	}
	for _, cfg := range config {
		go handle(cfg)
	}
}

func main() {

	d := sync.Map{}
	d.Store("1234", "hello")
	test1(d)
	load, _ := d.Load("1234")
	fmt.Println(load)
}

type Dmeo struct {
	name string
}

func test1(m sync.Map) {
	m.Store("1234", "1234")
}
