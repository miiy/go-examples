//1.下面的代码有什么问题？
package main

import "fmt"

type N int
func (n N) value(){
	n++
	fmt.Printf("v:%p,%v\n",&n,n)
}

func (n *N) pointer(){
	*n++
	fmt.Printf("v:%p,%v\n",n,*n)
}


func main() {

	var a N = 25

	p := &a
	p1 := &p

	p1.value()
	p1.pointer()
}

//参考答案及解析：编译错误：
//
//calling method value with receiver p1 (type **N) requires explicit dereference
//calling method pointer with receiver p1 (type **N) requires explicit dereference
//
//不能使用多级指针调用方法。
