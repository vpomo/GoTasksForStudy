package main

import (
	"fmt"
	"time"
)

type Result struct {
	Value int
	Err   error
}

func worker(id int, jobs <-chan int, results chan<- Result) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		if j%2 == 0 {
			results <- Result{Value: j * 2, Err: fmt.Errorf("error in job %d", j)}
		} else {
			results <- Result{Value: j * 2, Err: nil}
		}
		fmt.Printf("worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		result := <-results
		if result.Err != nil {
			fmt.Printf("Job %d failed: %v\n", result.Value/2, result.Err)
		} else {
			fmt.Printf("Job %d result: %d\n", result.Value/2, result.Value)
		}
	}
}
