package main

import (
	"fmt"
)

const XX = 2

func main() {
	const (
		XX = XX + XX
		Y
		Z
	)
	fmt.Println(XX, Y, Z) //4 8 8
}
