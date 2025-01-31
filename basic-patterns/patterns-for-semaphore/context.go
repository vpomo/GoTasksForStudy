package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerSc(ctx context.Context, id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case sem <- struct{}{}:
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
		<-sem
	case <-ctx.Done():
		fmt.Printf("Worker %d canceled\n", id)
	}
}

func main() {
	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerSc(ctx, i, sem, &wg)
	}

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancel")
		cancel()
	}()

	wg.Wait()
	fmt.Println("All workers done")
}
