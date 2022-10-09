package main

import "fmt"

func main() {
	fmt.Println(defer06())
}

func defer06() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//
// 1