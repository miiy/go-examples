package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a := [...]int{2,2,3,4}
	fmt.Println(reflect.TypeOf(a))
	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	for i,v := range []byte("世界abc") {
		fmt.Println(i, v)
	}
}
