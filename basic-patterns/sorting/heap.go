package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range heapSort(unsorted) {
		fmt.Println(num)
	}
}

func heapSort(numbers []int) []int {
	n := len(numbers)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(numbers, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		heapify(numbers, i, 0)
	}

	return numbers
}

func heapify(numbers []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && numbers[left] > numbers[largest] {
		largest = left
	}
	if right < n && numbers[right] > numbers[largest] {
		largest = right
	}

	if largest != i {
		numbers[i], numbers[largest] = numbers[largest], numbers[i]
		heapify(numbers, n, largest)
	}
}
