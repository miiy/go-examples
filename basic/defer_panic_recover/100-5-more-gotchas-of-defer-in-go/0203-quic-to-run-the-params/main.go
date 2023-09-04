package main

import "fmt"

//#3—Quick to run the params

// 传递给延迟函数的参数是在注册延迟时，而不是在运行时求值的。

// Example
type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}

func (p *message) print() string {
	return p.content
}

func main() {
	m := &message{
		content: "Hello",
	}
	defer fmt.Print(m.print())
	m.set("World")
	// deferred func runs
}

// Output:
// Hello

// 为什么输出的不是 "World"
// 在defer中，fmt.Print 被推迟直到函数返回，但 m.print() 会立即求值。因此 m.print() 将返回“Hello”，并将其保存在一边，直到包围的函数返回。
//
//func () {
//	m := &message{content: "Hello"}
//	defer fmt.Println(m.print()) // m.print() = &message{content: "Hello"}
//	m.set("World")
//	// deferred func runs: fmt.Print("Hello")
//}
