package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	once.Do(func() {
		fmt.Println("Only once")
	})

	once.Do(func() {
		fmt.Println(111)
	})
}

// output:
// Only once
