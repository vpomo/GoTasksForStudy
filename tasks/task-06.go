package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(ch chan int) {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 1; i < 10; i++ {
			ch <- r.Intn(i)
		}
		close(ch)
	}()
}

func main() {
	ch := make(chan int)
	go generator(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
