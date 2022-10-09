package main

import "fmt"

func main() {
	defer05()
}

func defer05() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//
// 0