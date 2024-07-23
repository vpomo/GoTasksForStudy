package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var pa *Student   // pa == nil
	pa = new(Student) // pa == &Student{"", 0}
	pa.Name = "Alex"  // pa == &Student{"Alex", 0}
	pa.Age = 22
	fmt.Println(pa)
}
