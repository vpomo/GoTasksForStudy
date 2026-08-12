package main

import (
	"fmt"
)

func main() {
	m := map[int64]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
	}

	for k := range m {
		fmt.Println(k)
	}
}

//1
//2
//3

//2
//3
//1
