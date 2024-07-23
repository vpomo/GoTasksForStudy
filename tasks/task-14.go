package main

import (
	"fmt"
)

func filter(done chan struct{}, inputStream <-chan int, operation func(int) bool) <-chan int {
	filteredStream := make(chan int)
	go func() {
		defer close(filteredStream)
		for {
			select {
			case <-done:
				return
			case num, ok := <-inputStream:
				if !ok {
					return
				}
				if !operation(num) {
					break
				}
				select {
				case <-done:
					return
				case filteredStream <- num:
				}
			}
		}
	}()

	return filteredStream
}

func main() {
	done := make(chan struct{})
	defer close(done)

	rangeNum := make(chan int)
	go func() {
		for _, num := range []int{1, 2, 3, 4, 7, 8, 9, 11, 12, 13, 14, 15} {
			rangeNum <- num
		}
		close(rangeNum)
	}()

	isEven := func(nm int) bool {
		return nm%2 == 0
	}

	for num := range filter(done, rangeNum, isEven) {
		fmt.Println(num)
	}
}
