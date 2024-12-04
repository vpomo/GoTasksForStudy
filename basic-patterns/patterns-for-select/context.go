package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c <- "result"
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("context timeout")
	}
}
