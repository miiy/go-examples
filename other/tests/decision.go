package main

import "fmt"

func main() {
	var a int  = 10
	var b int = 100

	// if
	if a < 20 {
		fmt.Println("a < 20")
	}

	// if else
	if b < 20 {
		fmt.Println("b < 20")
	} else {
		fmt.Println("b !< 20")
	}

	// if 嵌套
	if a == 10 {
		if b == 100 {
			fmt.Println("a = 10, b = 100")
		}
	}

	// switch

	// select
}

/*
$ go run decision.go 
a < 20
b !< 20
a = 10, b = 100
*/
