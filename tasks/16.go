package main

import (
	"fmt"
	"time"
)

// Функция, которая отправляет данные в канал
func sendData(ch chan<- int, start int, interval time.Duration) {
	for i := start; i < start+5; i++ {
		ch <- i
		time.Sleep(interval)
	}
	close(ch)
}

func main() {
	// Создаем несколько каналов
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go sendData(ch1, 1, 1*time.Second)
	go sendData(ch2, 10, 2*time.Second)
	go sendData(ch3, 100, 3*time.Second)

	// Объединяем данные из нескольких каналов в один
	for {
		select {
		case data, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println("Data from ch1:", data)
			}
		case data, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println("Data from ch2:", data)
			}
		case data, ok := <-ch3:
			if !ok {
				ch3 = nil
			} else {
				fmt.Println("Data from ch3:", data)
			}
		}

		// Если все каналы закрыты, выходим из цикла
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}

	fmt.Println("All channels are closed")
}

/*
Data from ch3: 100
Data from ch2: 10
Data from ch1: 1
Data from ch1: 2
Data from ch1: 3
Data from ch2: 11
Data from ch3: 101
Data from ch1: 4
Data from ch2: 12
Data from ch1: 5
Data from ch2: 13
Data from ch3: 102
Data from ch2: 14
Data from ch3: 103
Data from ch3: 104
All channels are closed
*/
