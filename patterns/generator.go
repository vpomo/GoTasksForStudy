package main

import "fmt"

// https://habr.com/ru/articles/852556/
func main() {
	// Данные, которые будут отправляться в канал
	items := []int{10, 20, 30, 40, 50}

	// Получаем канал с данными из генератора
	dataChannel := generator(items)

	// Потребитель обрабатывает данные из канала
	process(dataChannel)
}

// generator создает канал и запускает горутину для отправки данных
func generator(items []int) chan int {
	ch := make(chan int)

	go func() {
		// Закрываем канал после завершения отправки данных
		defer close(ch)

		// Перебираем элементы и отправляем их в канал
		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}

// process получает данные из канала и выводит их
func process(ch chan int) {
	for item := range ch {
		fmt.Println(item)
	}
}
