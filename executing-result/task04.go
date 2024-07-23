package main

import "fmt"

func main() {
	count := 0

	for i := range [256]struct{}{} {
		if n := byte(i); n == -n {
			count++
		}
	}
	fmt.Println(count) //2

	n := byte(128)
	fmt.Println(n)  //128
	fmt.Println(-n) //128
}
