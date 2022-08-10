package main

var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}

func main() {
	hello()
	// 使用select 造成死锁 ????

	select {}
}
