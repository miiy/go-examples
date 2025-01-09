package main

import "fmt"

func main()  {
	var a = 5
	fmt.Printf("%d %p \n", a, &a) // 5 0xc0000b4008
	var intP *int
	intP = &a
	fmt.Printf("%p", intP) // 0xc0000b4008
}