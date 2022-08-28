package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("start job %v %v", id, j)
		fmt.Println()
		time.Sleep(1 * time.Second)
		fmt.Println("end job %v %v ", id, j)
		result <- j
	}

}
func main() {
	const numjobs = 5
	jobs := make(chan int, numjobs)
	result := make(chan int, numjobs)

	for i := 1; i < 3; i++ {
		go worker(i, jobs, result)
	}
	for i := 0; i < numjobs; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 1; a < numjobs; a++ {
		<-result
	}
	fmt.Println(result)
}
