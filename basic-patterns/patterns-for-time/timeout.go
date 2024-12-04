package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c <- "result"
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

/**
Паттерн тайм-аута используется для ограничения времени выполнения операции.
*/
