package main

import (
	"fmt"
	"sync"
	"time"
)

func workerNwg(id int, wg *sync.WaitGroup, nestedWg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	nestedWg.Add(1)
	go func(nestedID int) {
		defer nestedWg.Done()
		fmt.Printf("Nested Worker %d starting\n", nestedID)
		time.Sleep(time.Second)
		fmt.Printf("Nested Worker %d done\n", nestedID)
	}(id)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	var nestedWg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerNwg(i, &wg, &nestedWg)
	}

	go func() {
		wg.Wait()
		nestedWg.Wait()
	}()

	fmt.Println("All workers and nested workers done")
}
