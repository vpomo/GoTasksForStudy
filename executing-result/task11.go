package main

import (
	"fmt"
)

type S struct {
	a, b, c string
}

type I struct {
	a, b, c int
}

type Student struct {
	Name string
	Age  int
}

func main() {
	x := &S{"a", "b", "c"}
	y := &S{"a", "b", "c"}
	fmt.Println(x, y)     //&{a b c} &{a b c}
	fmt.Println(*x, *y)   //{a b c} {a b c}
	fmt.Println(x == y)   //false
	fmt.Println(*x == *y) //true

	z := interface{}(&I{1, 2, 3})
	k := interface{}(&I{1, 2, 3})
	fmt.Println(z == k) //false

	d1 := Student{"Albert", 1}
	d2 := Student{"Albert", 1}
	fmt.Println(d1, d2)     // {Albert 1} {Albert 1}
	fmt.Println(d1 == d2)   // true
	fmt.Println(&d1 == &d2) // false

	e1 := &Student{"Albert", 1}
	e2 := &Student{"Albert", 1}
	fmt.Println(e1, e2)     // &{Albert 1} &{Albert 1}
	fmt.Println(e1 == e2)   // false
	fmt.Println(*e1 == *e2) // true

}
