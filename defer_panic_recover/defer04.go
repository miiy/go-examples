package main

import "fmt"

func main() {
	fmt.Println(defer04())
}

func defer04() int {
	defer fmt.Println(1)
	panic(2)
	return 0
}

// 1
// panic: 2
