package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	changes(s)
	fmt.Println(s)

	s2 := make([]int, 3, 10)
	s2 = append(s2, []int{1, 2, 3}...)

	fmt.Println(s2)
	changes(s)
	fmt.Println(s2)
}

func changes(s []int) {
	s = append(s, 4)
}
