package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 3 * time.Second
	startTime := time.Now().UnixNano()
	fmt.Println("start context ...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Println("Wait select")
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("After 3 sec")
	case <-time.After(2 * time.Second):
		fmt.Println("After 2 sec")
	case <-time.After(1 * time.Second):
		fmt.Println("After 1 sec")
	case <-ctx.Done():
		fmt.Println("selected done")
		fmt.Println(ctx.Err())
	}
	endTime := time.Now().UnixNano()
	fmt.Println(endTime - startTime)
}

/**
start context ...
Wait select
After 1 sec
1015404200
*/
