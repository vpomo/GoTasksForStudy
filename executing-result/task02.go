package main

import "fmt"

func f(n int) (r int) {
	a, r := n-1, n+1
	if a+a == r {
		c, r := n, n*n       //!!! this is new shadowed r
		fmt.Println("r=", r) //9
		r = r - c
		fmt.Println("r=", r) //6
	}
	return r //r==4
}

func main() {
	fmt.Println(f(3))
}

/*
r= 9
r= 6
4
*/
