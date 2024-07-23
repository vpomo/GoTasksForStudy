package main

import "fmt"

func main() {
	s1 := make([]int, 1)
	fmt.Println("s1", s1) // [0]
	s2 := s1
	s2[0] = 1
	fmt.Println("s1", s1) // [1]
	s1 = append(s1, 42)
	s2[0] = 21
	fmt.Println("s1", s1) // [1, 42]
	fmt.Println("s2", s2) // [21]
}
