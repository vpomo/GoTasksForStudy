package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
Вопрос: Какая из строк (A или B) позволит получить лучшее время выполнения при условии что у нас 8 ядер (numcpu=8)?
Oтвет: Строка B с одним процессором
*/

const MAX_STEP = 1000000

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("numCPU", numCPU)
	//runtime.GOMAXPROCS(numCPU) //A
	runtime.GOMAXPROCS(1) //B

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < MAX_STEP; i++ {
			ch1 <- i
		}
		ch1 <- -1
		ch2 <- 0.0
	}()

	go func() {
		total := 0.0
		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < MAX_STEP; i++ {
				m := <-ch1
				if m == -1 {
					ch2 <- total
				}
			}
			t2 := time.Now().UnixNano()
			diff := float64(t2-t1) / 1000000.0
			total += diff
			fmt.Println("diff", diff)
		}
	}()

	fmt.Println("Total: ", <-ch2, <-ch2) //Total:  0 164.0148
}
