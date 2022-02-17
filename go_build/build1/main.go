package main

import (
	"fmt"
)

func main()  {
	fmt.Println("hello.")
}

// go build main.go
// go build --o test_build main.go
//
// 交叉编译
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go