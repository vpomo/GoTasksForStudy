package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		mu      sync.Mutex
		cond    = sync.NewCond(&mu)
		ready   bool
		message string
	)

	// Горутина-ожидатель
	go func() {
		mu.Lock()
		defer mu.Unlock()

		// Ждем, пока ready не станет true
		fmt.Println("Ожидание сигнала ...")
		for !ready {
			cond.Wait() // Освобождает мьютекс и блокируется, пока не получит сигнал
		}

		fmt.Println("Ожидатель получил сигнал!")
		fmt.Println("Сообщение:", message)
	}()

	// Горутина-отправитель
	go func() {
		time.Sleep(2 * time.Second) // Имитация задержки

		mu.Lock()
		defer mu.Unlock()

		// Устанавливаем флаг и отправляем сигнал
		ready = true
		message = "Привет из cond!"
		cond.Signal() // Пробуждает одну ожидающую горутину
	}()

	// Даем время горутинам завершиться
	time.Sleep(3 * time.Second)
}
