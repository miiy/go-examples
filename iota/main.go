package main

import (
	"fmt"
)

const (
	mutexLocked      = 1 << iota // 1 << 0
	mutexWoken                   // 1 << 1
	mutexStarving                // 1 << 2
	mutexWaiterShift = iota      //      3
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

func main() {

	fmt.Println(IgEggs)
	fmt.Println(IgChocolate)
	fmt.Println(IgNuts)
	fmt.Println()
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
}
