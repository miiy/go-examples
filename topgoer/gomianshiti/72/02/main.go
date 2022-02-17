//2.下面的代码输出什么，请说明？
package main

import "fmt"

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

func main() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}

//参考答案及解析：{5}。这道题容易忽视的点是，String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法。
//
//可以这样修复：
//
//func main() {
//	orange := &Orange{}
//	orange.Increase(10)
//	orange.Decrease(5)
//	fmt.Println(orange)
//}
