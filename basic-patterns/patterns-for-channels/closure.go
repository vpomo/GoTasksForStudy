package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c <- "one"
		close(c)
	}()

	for {
		msg, ok := <-c
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("received", msg)
	}
}
