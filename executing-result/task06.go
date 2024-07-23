package main

import "fmt"

func main() {
	x := []int{123}
	fmt.Println(x)     //[123]
	x, x[0] = nil, 456 //[]
	fmt.Println(x)
	x, x[0] = []int{666}, 999 //panic: runtime error: index out of range [0] with length 0
	fmt.Println(x)
}
