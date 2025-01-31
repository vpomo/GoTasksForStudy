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
		time.Sleep(1 * time.Second)
		close(c)
	}()

	for {
		select {
		case msg, ok := <-c:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("received", msg)
		}
	}
}
