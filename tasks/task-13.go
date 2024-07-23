package main

import (
	"fmt"
	"time"
)

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	//s := fmt.Sprintf("Start worker № %d", id)
	//fmt.Println(s)
	for j := range jobs {
		results <- f(j)
		//s := fmt.Sprintf("worker № %d with job %d", id, j)
		//fmt.Println(s)
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 1; w <= 3; w++ {
		k := w
		go worker(k, multiplier, jobs, results)
	}

	var i = 0
	go func() {
		for {
			select {
			case num := <-results:
				s := fmt.Sprintf("Result work is %d", num)
				fmt.Println(s)
				i++
				if i == numJobs {
					fmt.Println("Finished")
					close(results)
					return
				}
			}
		}
	}()

	timer := time.NewTimer(time.Second)
	finished := <-timer.C
	fmt.Println(finished)
}
