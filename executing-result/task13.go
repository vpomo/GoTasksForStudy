package main

import "fmt"

func main() {
	s := "qwe"
	ps := &s
	fmt.Println("ps", ps) //0xc000024070

	b := []byte(*ps)
	fmt.Println("b", b) // [113 119 101]
	pb := &b
	fmt.Println("pb", pb) // &[113 119 101]

	s += "r"
	fmt.Println("s", s) // "qwer"
	*ps += "t"
	fmt.Println("*ps", *ps) // "qwert"
	*pb = append(*pb, []byte("y")[0])
	fmt.Println("*pb", *pb) // [113 119 101 121] == "qwey"
	fmt.Println("*ps", *ps) // "qwert"
}
