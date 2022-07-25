package main

import "fmt"

func main() {
	fmt.Println(defer07())
}

func defer07() int {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

//
// 5