package main

import "fmt"

type A struct {
	g int
}

func (A) m() int {
	return 1
}

type B int

func (B) g() {}
func (B) f() {}

type C struct {
	A
	B
}

func (C) m() int {
	return 9
}

func main() {
	var c interface{} = C{}
	fmt.Println(c) // {{0} 0}

	a, bf := c.(interface{ f() })
	b, bg := c.(interface{ g() })
	i := c.(interface{ m() int })
	fmt.Println(a, b)          // {{0} 0} <nil>
	fmt.Println(i)             // {{0} 0}
	fmt.Println(bf, bg, i.m()) // true false 9
}
