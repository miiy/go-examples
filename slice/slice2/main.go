package main

import (
	"fmt"
)

func main()  {
	s := "在"
	fmt.Println(len(s))


	a := []int{1, 2, 3}
	fmt.Println(a)
	changeSlice(a)
	fmt.Println(a)

	slice2()
}

func changeSlice (s []int)  {
	s[1] = 2
}

func slice2()  {
	tests := []int{1,2,3,4,5}
	test2 := tests[2:4]
	fmt.Println(tests)
	fmt.Println(test2)
	test2[0] = 0
	fmt.Println(tests)
	fmt.Println(test2)
}