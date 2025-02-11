package main

import "fmt"

func main() {
	unsorted := []int{2, 56, 1, 8, 32}
	for _, num := range mergeSort(unsorted) {
		fmt.Println(num)
	}
}

func mergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
