package main

import "fmt"

// #5 — Deferred method gotchas
// #5 — Deferred method 陷阱

// Without pointers
type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Println(c.model)
}

func main() {
	c := Car{model: "DeLorean DMC-12"}
	defer c.PrintModel()

	c.model = "Chevrolet Impala"
}

// Output:
// DeLorean DMC-12

// WithPointers
//
//func (c *Car) PrintModel() {
//	fmt.Println(c.model)
//}

// Output:
// Chevrolet Impala

// What's going on?
//
// func (receiver Type) f(input) result
// func f(receiver Type, input) result
//
// 请记住，传递给 deferred func 的参数将立即保到一边，而无需等待延迟函数运行
// 因此，当带有值接收器的方法与 defer 一起使用时，接收器将在注册时被拷贝 (in this case Car)，并且对其所做的修改将不可见 (Car.model)。
// 因为接收器也是一个输入参数，并且 defer 注册时立即求值为 ”DeLorean DMC-12“。
// 另一方面，当接收器是指针时，当它被 defer 调用时，会创建一个新的指针，但它指向的地址将与上面的“c”指针相同。因此，对它的任何修改都将完美地反映出来。
