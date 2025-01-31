package main

import "fmt"

func intersection(a, b []int) []int {
	result := []int{}
	counter := make(map[int]int)

	for _, num := range a {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num] += 1
		}
	}

	for _, num := range b {
		if count, ok := counter[num]; ok && count > 0 {
			counter[num] -= 1
			result = append(result, num)
		}
	}

	return result
}

func main() {

	a := []int{37, 5, 1, 2, 8, 8}
	b := []int{6, 8, 2, 4, 37}
	// [8, 2, 37]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}
