package main

import "fmt"

func main() {
	defer02()
	// defer02_01()
}

func defer02() {
	defer fmt.Println(1)
	defer func() {
		panic(2)
	}()
	defer fmt.Println(3)
}

func defer02_01() {
	defer fmt.Println(1)
	defer panic(2)
	defer fmt.Println(3)
}

//
// 3
// 1
// panic: 2
