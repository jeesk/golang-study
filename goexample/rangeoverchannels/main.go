package main

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)
	// 没有关闭的话是不能使用 普通的for 循环

}
