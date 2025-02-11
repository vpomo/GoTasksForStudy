package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range selectionSort(unsorted) {
		fmt.Println(num)
	}
}

func selectionSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}

	return numbers
}
