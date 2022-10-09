package main

import "fmt"

func main() {
	fmt.Println(defer08())
}

func defer08() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//
// 5