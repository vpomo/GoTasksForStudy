package main

import "fmt"

func main() {
	var x interface{} = []int{1, 2, 3}
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType) // "[]int"

	var y interface{} = 2.3
	switch v := y.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
	// Вывод: float64: 2.3
}
