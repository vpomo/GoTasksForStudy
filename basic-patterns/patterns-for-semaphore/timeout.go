package main

import (
	"fmt"
	"sync"
	"time"
)

func workerSt(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case sem <- struct{}{}:
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
		<-sem
	case <-time.After(2 * time.Second):
		fmt.Printf("Worker %d timed out\n", id)
	}
}

func main() {
	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerSt(i, sem, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
