package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

// Output:
// 3
// 2
// 1

// defer的顺序是先进后出
