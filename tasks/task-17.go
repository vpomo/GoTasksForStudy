package main

import (
	"fmt"
	"time"
)

func mergeC(channels ...<-chan int) <-chan int {

	result := make(chan int)
	for _, ch := range channels {
		go func(channel <-chan int) {
			for num := range channel {
				result <- num
			}
		}(ch)
	}

	go func() {
		for range channels {
			<-time.After(time.Second) // Даем время для завершения горутин
		}
		close(result)
	}()

	return result
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 1; i < 4; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1)
	}()

	go func() {
		for i := 4; i < 7; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
		close(ch2)
	}()

	go func() {
		for i := 7; i < 10; i++ {
			ch3 <- i
			time.Sleep(time.Second)
		}
		close(ch3)
	}()

	for num := range mergeC(ch1, ch2, ch3) {
		fmt.Println(num)
	}
	fmt.Println("Finished")
}
