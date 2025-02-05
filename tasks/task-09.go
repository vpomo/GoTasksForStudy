package main

import (
	"fmt"
	"sync"
)

func readChan(a []int, wg *sync.WaitGroup, outCh chan int) {
	defer wg.Done()
	for _, num := range a {
		outCh <- num
	}
}

func mergeSlice(x, y []int, ch chan int) {
	defer close(ch)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go readChan(x, &wg, ch)
	go readChan(y, &wg, ch)

	wg.Wait()
}

func main() {
	//Case 1
	a := []int{10, 20, 30}
	b := []int{40, 50, 60}
	h := []int{}

	ch := make(chan int)
	go mergeSlice(a, b, ch)

	for num := range ch {
		h = append(h, num)
	}
	fmt.Println(h)

	//Case 2
	c := []int{1, 2, 3}
	d := []int{4, 5, 6}
	e := append(c, d...)
	fmt.Println(e)
}

// [10 20 40 50 60 30]
// [1 2 3 4 5 6]
