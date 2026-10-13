package main

import "fmt"

// Измените порядок целых чисел в слайсе.

func reverse1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s []int) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	a := []int{2, 1, 4, 6, 3, 8}
	b := []int{2, 1, 4, 6, 3, 8}

	reverse1(a)
	fmt.Println(a) //[8 3 6 4 1 2]

	reverse2(b)
	fmt.Println(b) //[8 3 6 4 1 2]
}
