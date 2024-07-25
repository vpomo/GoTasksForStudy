package main

import (
	"fmt"
	"time"
)

func worker01() chan int {
	ch := make(chan int)

	go func() {
		fmt.Println("before sleeping")
		time.Sleep(time.Second * 3)
		ch <- 42
		fmt.Println("wrote down to channel")
	}()

	return ch
}
func main() {
	fmt.Println("reading ...")
	_, _ = <-worker01(), <-worker01()
	fmt.Println("finished")
}

/**
За сколько времени выполнится приложение?
*/

/**
6 sec
reading ...
before sleeping
write down to channel
before sleeping
write down to channel
finished
*/
