package main

import "fmt"

func main() {
	chInput := make(chan int)
	chSquare := make(chan int)
	go func() {
		numbers := []int{1, 3, 4, 5}
		for _, num := range numbers {
			chInput <- num
		}
		close(chInput)
	}()
	go func() {
		for num := range chInput {
			chSquare <- num * num
		}
		close(chSquare)
	}()

	for num := range chSquare {
		fmt.Println(num)
	}
}
