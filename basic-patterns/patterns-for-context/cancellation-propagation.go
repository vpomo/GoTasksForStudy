package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	childCtx, childCancel := context.WithCancel(ctx)
	defer childCancel()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	select {
	case <-childCtx.Done():
		fmt.Println("Child context canceled")
	}
}
