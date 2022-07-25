package main

import "fmt"

func main() {
	fmt.Println(defer03())
}

func defer03() int {
	defer fmt.Println(1)
	return 2
}

//
// 1
// 2