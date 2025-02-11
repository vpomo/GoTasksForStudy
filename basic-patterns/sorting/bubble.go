package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range bubbleSort(unsorted) {
		fmt.Println(num)
	}
}

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return numbers
}
