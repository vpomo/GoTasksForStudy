package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var pa *Student // pa == nil
	fmt.Println(pa) // <nil>

	pa = new(Student) // pa == &Student{"", 0}
	fmt.Println(pa)   // &{ 0}

	pa.Name = "Alex" // pa == &Student{"Alex", 0}
	fmt.Println(pa)  // &{Alex 0}

	pa.Age = 22
	fmt.Println(pa) // &{Alex 22}
}
