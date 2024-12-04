package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Priority int
}

func worker(id int, sem chan struct{}, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		select {
		case sem <- struct{}{}:
			fmt.Printf("Worker %d starting task %d with priority %d\n", id, task.ID, task.Priority)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done task %d\n", id, task.ID)
			<-sem
		}
	}
}

func main() {
	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup
	tasks := make(chan Task, 5)

	for w := 1; w <= maxWorkers; w++ {
		wg.Add(1)
		go worker(w, sem, tasks, &wg)
	}

	for i := 1; i <= 5; i++ {
		tasks <- Task{ID: i, Priority: i % 3}
	}
	close(tasks)

	wg.Wait()
	fmt.Println("All workers done")
}
