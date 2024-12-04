package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")

	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("Value from context:", ctx.Value("key"))
	}(ctx)

	time.Sleep(2 * time.Second)
}
