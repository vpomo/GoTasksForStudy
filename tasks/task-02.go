package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{2, 1, 4, 6, 3, 8}
	reverse(a)
	fmt.Println(a) //[8 3 6 4 1 2]
}
