package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := 10
	fmt.Println("a")
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &a)

	b := [1]int{}
	fmt.Println("b")
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &b)

	c := [2]int{1}
	fmt.Println("c")
	fmt.Printf("%p\n", c)
	fmt.Printf("%p\n", &c)

	d := []int{1}
	fmt.Println("d")
	fmt.Printf("%p\n", d)
	fmt.Printf("%p\n", (*reflect.SliceHeader)(unsafe.Pointer(&d)))
	fmt.Printf("%p\n", &d)

}
