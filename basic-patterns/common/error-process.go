package main

import (
	"errors"
	"log"
)

// ResultErr — структура для хранения результата и ошибки
type ResultErr struct {
	data int
	err  error
}

func main() {
	input := []int{1, 2, 3, 4}

	resultCh := make(chan ResultErr)

	// запускаем потребителя, который будет отправлять результаты и ошибки
	go consumer(generatorErr(input), resultCh)

	// читаем результаты
	for res := range resultCh {
		if res.err != nil {
			log.Println("Ошибка:", res.err)
		} else {
			log.Println("Результат:", res.data)
		}
	}
}

// generatorErr отправляет данные в канал
func generatorErr(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)
		for _, data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}

// consumer вызывает функцию, которая может возвращать ошибку
func consumer(inputCh chan int, resultCh chan ResultErr) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := callDatabase(data)
		resultCh <- ResultErr{data: resp, err: err}
	}
}

// callDatabase возвращает ошибку
func callDatabase(data int) (int, error) {
	if data == 3 {
		return data, errors.New("ошибка запроса к базе данных")
	}
	return data, nil
}
