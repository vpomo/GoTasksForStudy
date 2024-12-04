package main

import (
	"fmt"
	"time"
)

type ResultS struct {
	Value int
	Err   error
}

func workerEs(id int, results chan<- ResultS) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	if id%2 == 0 {
		results <- ResultS{Value: id * 2, Err: fmt.Errorf("error in worker %d", id)}
	} else {
		results <- ResultS{Value: id * 2, Err: nil}
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	results := make(chan ResultS, 3)

	for i := 1; i <= 3; i++ {
		go workerEs(i, results)
	}

	for i := 0; i < 3; i++ {
		select {
		case result := <-results:
			if result.Err != nil {
				fmt.Printf("Worker %d failed: %v\n", result.Value/2, result.Err)
			} else {
				fmt.Printf("Worker %d result: %d\n", result.Value/2, result.Value)
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout waiting for results")
		}
	}
}
