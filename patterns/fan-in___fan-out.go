package main

import (
	"fmt"
	"sync"
	"time"
)

// https://habr.com/ru/articles/852556/
// generatorFan — создает канал с данными
func generatorFan(doneCh chan struct{}, numbers []int) chan int {
	dataStream := make(chan int)

	go func() {
		defer close(dataStream)
		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case dataStream <- num:
			}
		}
	}()

	return dataStream
}

// addFan — добавляет 1 к каждому значению
func addFan(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)
		for num := range inputCh {
			// Имитация более затратной работы
			time.Sleep(time.Second)
			result := num + 1

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

// multiplyFan — умножает каждое значение на 2
func multiplyFan(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)
		for num := range inputCh {
			result := num * 2

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

// fanOut — создает несколько горутин addFan для параллельной обработки данных
func fanOut(doneCh chan struct{}, inputCh chan int, workers int) []chan int {
	resultChannels := make([]chan int, workers)

	for i := 0; i < workers; i++ {
		resultChannels[i] = addFan(doneCh, inputCh)
	}

	return resultChannels
}

// fanIn — объединяет результаты нескольких каналов в один
func fanIn(doneCh chan struct{}, channels ...chan int) chan int {
	finalStream := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		chCopy := ch
		wg.Add(1)

		go func() {
			defer wg.Done()
			for value := range chCopy {
				select {
				case <-doneCh:
					return
				case finalStream <- value:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalStream)
	}()

	return finalStream
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generatorFan(doneCh, numbers)

	// создаем 10 горутин add с помощью fanOut
	channels := fanOut(doneCh, inputCh, 10)

	// объединяем результаты из всех каналов
	addResultCh := fanIn(doneCh, channels...)

	// передаем результат в следующий этап multiply
	resultCh := multiplyFan(doneCh, addResultCh)

	// выводим результаты
	for result := range resultCh {
		fmt.Println(result)
	}
}
