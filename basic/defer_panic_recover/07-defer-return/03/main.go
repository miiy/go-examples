package main

import "fmt"

func test() (r int) {
	defer func(r int) {
		fmt.Println("defer:", r)
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println("return:", test())
}

// Output:
// defer: 0
// return: 1
