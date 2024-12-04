package main

import (
	"fmt"
	"sync"
	"time"
)

func workerRc(id int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	results <- id * 2
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerRc(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("All workers done")
}
