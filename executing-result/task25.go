package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	var s1 []int         // nil срез, Data = 0, Len = 0, Cap = 0
	s2 := make([]int, 0) // пустой срез (не nil),Data != 0 (указывает на массив), Len = 0, Cap = 0

	sr1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	// &s1 — берем указатель на срез
	// unsafe.Pointer(&s1) — преобразуем в указатель произвольного типа
	// (*reflect.SliceHeader)(...) — интерпретируем память как структуру SliceHeader
	// type SliceHeader struct {
	//	Data uintptr  // указатель на базовый массив
	//	Len  int      // длина среза
	//	Cap  int      // емкость среза
	// }

	sr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))

	fmt.Println(*sr1) // {0 0 0}
	fmt.Println(*sr2) // {824634064616 0 0}
	time.Sleep(time.Second)
}

// Process finished with the exit code 0
