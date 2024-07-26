package main

import (
	"fmt"
)

func main() {
	var m map[string]float64
	//m["pi"] = 3.141628 //panic: assignment to entry in nil map

	var n = make(map[string]float64)
	n["pi"] = 3.141628

	fmt.Println("m", m) //map[]
	fmt.Println("n", n) //map[pi:3.141628]
}
