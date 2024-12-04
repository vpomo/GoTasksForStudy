package main

import (
	"fmt"
	"sync"
	"time"
)

func workerSb(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{} // Захватываем семафор
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	<-sem // Освобождаем семафор
}

func main() {
	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerSb(i, sem, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
