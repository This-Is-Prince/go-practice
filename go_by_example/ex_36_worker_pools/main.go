package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, "started job:", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, "finished job:", j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("-----Worker Pools-----")

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Println("result is", result)
	}
}
