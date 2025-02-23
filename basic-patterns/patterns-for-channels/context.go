package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func workerC(ctx context.Context, c chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case c <- rand.Intn(10):
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan int)
	defer close(c)

	go workerC(ctx, c)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	cancel()
}
