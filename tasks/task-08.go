package main

import "fmt"

func main() {
	testMap := make(map[string]int)
	testMap["one"] = 100

	key := "two"
	fmt.Println("For key=two")
	if val, ok := testMap[key]; !ok {
		fmt.Printf("There isn't key and val=%d", val)
		fmt.Println("")
	}

	key = "one"
	fmt.Println("For key=one")
	if val, ok := testMap[key]; ok {
		fmt.Printf("There is key and val=%d", val)
	}
}
