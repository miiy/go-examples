package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界。")
}

/*
# 使用 go run 命令执行 Go 语言代码
$ go run hello.go 
Hello, world.
*/

/*
# 使用 go build 命令编译生成二进制文件：
$ go build hello.go
$ ls
hello  hello.go
$ ./hello 
Hello, world.
*/

// Go语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。