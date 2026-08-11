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

	for len(question) > 0 {
		c, size = utf8.DecodeRuneInString(question)
		fmt.Printf("c: %c %v bytes \n", c, size)
		question = question[size:]
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

c: ¿ 2 bytes
c: C 1 bytes
c: ó 2 bytes
c: m 1 bytes
c: o 1 bytes
c:   1 bytes
c: e 1 bytes
c: s 1 bytes
c: t 1 bytes
c: á 2 bytes
c: s 1 bytes
c: ? 1 bytes
*/
