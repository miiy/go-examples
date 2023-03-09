package main

import (
	"fmt"
)

func main() {
	s := "在"
	fmt.Println(len(s)) // 3

	tests := []int{1, 2, 3, 4, 5}
	test2 := tests[2:4]
	fmt.Println(tests) // [1, 2, 3, 4, 5]
	fmt.Println(test2) // [3, 4]
	test2[0] = 0
	fmt.Println(tests) // [1, 2, 0, 4, 5]
	fmt.Println(test2) // [0, 4]
}
