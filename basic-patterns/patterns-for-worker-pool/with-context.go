package main

import (
	"context"
	"fmt"
	"time"
)

func workerCtx(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case j := <-jobs:
			fmt.Printf("worker %d started job %d\n", id, j)
			time.Sleep(time.Second)
			fmt.Printf("worker %d finished job %d\n", id, j)
			results <- j * 2
		}
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for w := 1; w <= 3; w++ {
		go workerCtx(ctx, w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	cancel()
	time.Sleep(time.Second) // Даем время goroutines завершиться
}
