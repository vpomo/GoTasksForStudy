package main

import (
	"fmt"
)

type SStr struct {
	m string
}

func fHello() *SStr {
	return &SStr{"hello"}
}

func main() {
	p := fHello()
	fmt.Println(p.m) //hello
}
