package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes \n", c, size) // Выводит: First rune: ¿ 2 bytes
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}

// First rune: ¿ 2 bytes
/*
0 ¿
2 C
3 ó
5 m
6 o
7
8 e
9 s
10 t
11 á
13 s
14 ?
*/
