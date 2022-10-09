package main

import "fmt"

// 字符串拼接
func main() {
	s := "hel" + "lo, "
	s += "world!"
	fmt.Println(s) // hello, world!
}
