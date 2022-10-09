package main

import "fmt"

func main()  {
	// 空值赋值
	s := []int{}
	s = append(s, 1)
	fmt.Println(s) // [1]

	m := map[string]int{}
	m["test"] = 1
	fmt.Println(m) // map[test:1]

	// nil值赋值
	var s2 []int
	s2 = append(s2, 1)
	fmt.Println(s2) // [1]

	var m2 map[string]int
	m2["test"] = 1 // panic: assignment to entry in nil map
	fmt.Println(m2)
}