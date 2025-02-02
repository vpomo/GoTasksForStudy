package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Программа стартовала")
	// Создаем мьютекс и условную переменную
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// Общее условие
	conditionMet := false

	// Горутина, которая будет ждать условия
	go func() {
		mu.Lock()
		for !conditionMet {
			cond.Wait()
		}
		fmt.Println("Условие выполнено!")
		mu.Unlock()
	}()

	// Горутина, которая устанавливает условие
	go func() {
		time.Sleep(2 * time.Second) // Имитация задержки
		mu.Lock()
		conditionMet = true
		cond.Signal() // Сигнализируем, что условие выполнено
		mu.Unlock()
	}()

	// Основная горутина ждет завершения
	fmt.Println("Ждем завершения")
	time.Sleep(3 * time.Second)
	fmt.Println("Программа завершена")
}
