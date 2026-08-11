package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d canceled\n", id)
			return
		default:
			fmt.Printf("Worker %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

/*
Worker 3 working
Worker 2 working
Worker 1 working
Worker 2 working
Worker 3 working
Worker 1 working
Worker 1 working
Worker 3 working
Worker 2 working
Worker 3 working
Worker 1 working
Worker 2 working
Worker 1 canceled
Worker 2 canceled
Worker 3 canceled
*/
