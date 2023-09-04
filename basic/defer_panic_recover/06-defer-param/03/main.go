package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}

// case1 output:
//4
//3
//2
//1
//0

func case2() {
	for i := 0; i < 5; i++ {
		defer func() {
			println(i)
		}()
	}
}

// case2 output:
//5
//5
//5
//5
//5
