package main

import (
	"fmt"
	"sync"
	"time"
)

func workerEh(id int, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	if id%2 == 0 {
		errCh <- fmt.Errorf("error in worker %d", id)
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerEh(i, &wg, errCh)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("All workers done")
}
