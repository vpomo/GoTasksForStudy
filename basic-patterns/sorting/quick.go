package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range quickSort(unsorted) {
		fmt.Println(num)
	}
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	pivot := numbers[len(numbers)/2]
	left := []int{}
	right := []int{}

	for _, num := range numbers {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...) // Учитываем равные pivot элементы
}
