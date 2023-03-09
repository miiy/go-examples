package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1) // []
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(len(s1), s1) // 10 [0 1 2 3 4 5 6 7 8 9]

	s2 := make([]int, 0)
	fmt.Println(s2) // []
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
	}
	fmt.Println(len(s2), s2) // 10 [0 1 2 3 4 5 6 7 8 9]

	s3 := make([]int, 1)
	fmt.Println(s3) // [0]
}
