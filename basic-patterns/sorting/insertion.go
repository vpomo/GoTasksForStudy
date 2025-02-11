package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range insertionSort(unsorted) {
		fmt.Println(num)
	}
}

func insertionSort(numbers []int) []int {
	n := len(numbers)
	for i := 1; i < n; i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = key
	}
	return numbers
}
