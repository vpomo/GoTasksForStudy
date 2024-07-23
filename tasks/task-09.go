package main

import (
	"fmt"
	"sync"
)

func mergeSlice(x, y []int, ch chan int) {
	defer close(ch)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(a []int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, num := range a {
			ch <- num
		}
	}(x, &wg)
	go func(a []int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, num := range a {
			ch <- num
		}
	}(y, &wg)

	wg.Wait()
}

func main() {
	//Case 1
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
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
