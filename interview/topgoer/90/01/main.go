//1.下面代码能通过编译吗？
package main

type T int

func F(t T) {}

func main() {
	var q int
	F(q)
}