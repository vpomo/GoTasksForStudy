package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	ID  int
	Err error
}

func workerSeh(id int, sem chan struct{}, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case sem <- struct{}{}:
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		if id%2 == 0 {
			results <- Result{ID: id, Err: fmt.Errorf("error in worker %d", id)}
		} else {
			results <- Result{ID: id, Err: nil}
		}
		fmt.Printf("Worker %d done\n", id)
		<-sem
	}
}

func main() {
	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup
	results := make(chan Result, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerSeh(i, sem, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Worker %d failed: %v\n", result.ID, result.Err)
		} else {
			fmt.Printf("Worker %d succeeded\n", result.ID)
		}
	}

	fmt.Println("All workers done")
}
