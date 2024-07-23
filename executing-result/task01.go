package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	fmt.Println(a) //[0 1 2 3]
	x := a[:1]
	fmt.Println(x) //[0]
	y := a[2:]
	fmt.Println(y) //[2 3]
	x = append(x, y...)
	fmt.Println(x) // [0 2 3]
	fmt.Println(a) // [0 2 3 3]
	x = append(x, y...)
	fmt.Println(a, x) // [0 2 3 3] [0 2 3 3 3]
}
