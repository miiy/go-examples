package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
# go run package.go
My favorite number is 1
*/

// 每个go程序都是由包组成的