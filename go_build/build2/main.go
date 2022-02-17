package main

import "fmt"
import "github.com/miiy/go-examples/go_build/build2/config"

func main()  {
	fmt.Println("this build is:", config.String)
}

// https://zhuanlan.zhihu.com/p/92235251
//
// $ go build -tags dev main.go
// $ ./main
// this build is: this is debug mode
// $ go build -tags prod main.go
// $ ./main
// this build is: this is prod mode

// idea setting
//
// Languages & Frameworks -> Go -> Build Tags & Vendoring
// Custom tags: dev
//
// Run
// Check Use all custom build tags