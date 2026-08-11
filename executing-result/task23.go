package main

import (
	"fmt"
)

func main() {
	var a []int      // len=0, cap=0, указатель=nil
	a = append(a, 1) // len=1, cap=1
	a = append(a, 2) // len=2, cap=2 (перевыделение)
	a = append(a, 3) // len=3, cap=4 (перевыделение)

	b := append(a, 4)
	c := append(a, 5)

	fmt.Println("b:", b) //b: [1 2 3 5]
	fmt.Println("c:", c) //c: [1 2 3 5]
}

/**
Изначально: a -> [1, 2, 3, 0] (cap=4)
После b:   a -> [1, 2, 3, 4]
           b -> [1, 2, 3, 4]

После c:   a -> [1, 2, 3, 5]
           b -> [1, 2, 3, 5]
           c -> [1, 2, 3, 5]

Как исправить, чтобы получить ожидаемый результат?
// Способ 1: Использовать copy
b := make([]int, len(a))
copy(b, a)
b = append(b, 4)

c := make([]int, len(a))
copy(c, a)
c = append(c, 5)

// Способ 2: Использовать полное выражение среза для ограничения cap
b := append(a[:len(a):len(a)], 4)  // cap = len, принудительное выделение
c := append(a[:len(a):len(a)], 5)
*/
