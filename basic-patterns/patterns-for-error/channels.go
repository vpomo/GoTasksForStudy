package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value int
	Err   error
}

func workerEch(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, j)
		time.Sleep(time.Second)
		if j%2 == 0 {
			results <- Result{Value: j * 2, Err: fmt.Errorf("error in job %d", j)}
		} else {
			results <- Result{Value: j * 2, Err: nil}
		}
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerEch(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Job %d failed: %v\n", result.Value/2, result.Err)
		} else {
			fmt.Printf("Job %d result: %d\n", result.Value/2, result.Value)
		}
	}
}
