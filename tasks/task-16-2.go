package main

import (
	"fmt"
	"time"
)

// mergeChannels объединяет данные из нескольких каналов в один
func mergeChannels(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	// Запускаем горутину для каждого входного канала
	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				merged <- v
			}
		}(ch)
	}

	// Закрываем объединенный канал, когда все входные каналы закрыты
	go func() {
		for range channels {
			<-time.After(100 * time.Millisecond) // Даем время для завершения горутин
		}
		close(merged)
	}()

	return merged
}

func main() {
	// Создаем несколько каналов
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
			time.Sleep(150 * time.Millisecond)
		}
	}()

	go func() {
		defer close(ch3)
		for i := 7; i <= 9; i++ {
			ch3 <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Объединяем каналы
	merged := mergeChannels(ch1, ch2, ch3)

	// Читаем данные из объединенного канала
	for v := range merged {
		fmt.Println(v)
	}
}
