package main

import "fmt"

var f = func(x int) {
	fmt.Print("@@@")
}

func Bar() {
	f := func(x int) {
		if x >= 0 {
			fmt.Print(x) //2
			f(x - 1)
		}
	}
	f(2)
}

func Foo() {
	f = func(x int) {
		if x >= 0 {
			fmt.Print(x)
			f(x - 1)
		}
	}
	f(2)
}

func main() {
	Bar()
	fmt.Print("|")
	Foo()
}

//2@@@|210
