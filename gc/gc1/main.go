package main

import (
	"fmt"
)

func main()  {
	add()
	// 返回指针
	add2()
	// interface 类型
	// func Println(a ...interface{}) (n int, err error)
	fmt.Println("a")
	a := "a"
	fmt.Println(a)
}


func add() int {
	i := 10
	return i
}

func add2() *int {
	i := 10
	return &i
}

// $ go build -gcflags="-m -l" main.go
// # command-line-arguments
// ./main.go:25:2: moved to heap: i
// ./main.go:13:13: ... argument does not escape
// ./main.go:13:14: "a" escapes to heap
// ./main.go:15:13: ... argument does not escape
// ./main.go:15:13: a escapes to heap
