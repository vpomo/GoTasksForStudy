package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	oldStr := "\xbdm⌘fgah134517095aldrfgvh8h"
	var newStr strings.Builder

	for i := 0; i < len(oldStr); i++ {
		r, _ := utf8.DecodeRuneInString(string(oldStr[i]))
		fmt.Println(string(r))
		if unicode.IsNumber(r) {
			newStr.WriteString(string(r))
		}
	}
	fmt.Println(newStr.String())

	//work with index of slice
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[:1])
	fmt.Println(a[:0])
	fmt.Println(a[0:2])
	fmt.Println(a[1:2])
	fmt.Println(a[2:2])
	fmt.Println(a[2:3])
	fmt.Println(a[4:5])
}
