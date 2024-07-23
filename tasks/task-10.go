package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"

	if i, err := strconv.Atoi(s); err == nil {
		fmt.Printf("i=%d,type: %T\n", i, i) // i=123,type: int
	}

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		fmt.Printf("i=%d,type: %T\n", i, i) // i=123,type: int64
	}
}
