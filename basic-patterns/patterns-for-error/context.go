package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerEc(ctx context.Context, id int, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		errCh <- ctx.Err()
		return
	default:
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		if id%2 == 0 {
			errCh <- fmt.Errorf("error in worker %d", id)
		}
		fmt.Printf("Worker %d done\n", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerEc(ctx, i, &wg, errCh)
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
