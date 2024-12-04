package main

import (
	"fmt"
	"time"
)

func heartbeat(interval time.Duration, beat chan<- time.Time) {
	for {
		beat <- time.Now()
		time.Sleep(interval)
	}
}

func main() {
	beat := make(chan time.Time)
	go heartbeat(1*time.Second, beat)

	for {
		select {
		case b := <-beat:
			fmt.Println("Heartbeat at", b)
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

/*
Паттерн сердцебиения используется для периодической проверки состояния системы или подсистемы.
*/
