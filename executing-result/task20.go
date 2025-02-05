package main

import (
	"fmt"
)

type SS struct {
	name string
}

func main() {
	m := map[string]SS{"x": SS{"one"}}
	m["x"] = SS{name: "two"}
	fmt.Println(m["x"].name) //two

	n := map[string]SS{"x": SS{"one"}}
	fmt.Println(n["x"].name) //one
}
