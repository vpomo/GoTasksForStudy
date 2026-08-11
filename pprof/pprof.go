package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

// Функция, создающая CPU нагрузку
func cpuIntensiveWork() {
	for i := 0; i < 100000; i++ {
		for j := 0; j < 1000; j++ {
			_ = i * j
		}
	}
}

// Функция, работающая с памятью
func memoryIntensiveWork() {
	slice := make([][]byte, 0)
	for i := 0; i < 1000; i++ {
		data := make([]byte, 1024*1024) // 1MB на итерацию
		slice = append(slice, data)
	}
	time.Sleep(time.Second)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		cpuIntensiveWork()
	}()

	go func() {
		defer wg.Done()
		memoryIntensiveWork()
	}()

	wg.Wait()
	fmt.Fprintf(w, "Work completed at %s\n", time.Now().Format(time.RFC3339))
}

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	http.HandleFunc("/", handler)
	log.Println("Server starting on :8080")
	log.Println("Profile available at: http://localhost:8080/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
