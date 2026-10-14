package main

import (
	"context"
	"sync"
	"time"
)

type Action func(ctx context.Context)

func processConcurrently(ctx context.Context, action Action, n int) {
	if n <= 0 {
		return
	}

	// Семафор для ограничения конкурентности: разрешаем одновременно не более n задач
	sem := make(chan struct{}, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Захватываем слот в семафоре
			select {
			case sem <- struct{}{}:
				// Слот захвачен — запускаем задачу
				action(ctx)
				<-sem // Освобождаем слот после завершения (даже если паника не случится, defer тут не нужен, т.к. это не defer)
			case <-ctx.Done():
				// Контекст уже отменён до захвата слота — ничего не делаем
				return
			}
		}()
	}

	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	action := func(ctx context.Context) {
		// Имитация работы с проверкой ctx
		for {
			select {
			case <-time.After(500 * time.Millisecond):
				// Продолжаем работу
			case <-ctx.Done():
				return // Корректно останавливаемся при отмене
			}
		}
	}

	processConcurrently(ctx, action, 3)
}
