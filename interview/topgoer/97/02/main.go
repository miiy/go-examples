//2.下面代码输出什么？请简要说明。
package main

import "fmt"

type Foo struct {
	val int
}

func (f Foo) Inc(inc int) {
	f.val += inc
}

func main()  {
	var f Foo
	f.Inc(100)
	fmt.Println(f.val)
}

//参考答案及解析：输出 0。使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响的原来的值。如果想要在调用函数中修改原值，可以使用指针接收者定义的方法。
//
//type Foo struct {
//	val int
//}
//
//func (f *Foo) Inc(inc int) {
//	f.val += inc
//}
//
//func main() {
//	f := &Foo{}
//	f.Inc(100)
//	fmt.Println(f.val)  // 100
//}