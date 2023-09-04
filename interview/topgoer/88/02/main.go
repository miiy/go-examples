//2.下面的代码输出什么，请简要说明？
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	// ...
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

//参考答案及解析：nil false。知识点：接口值与 nil 值。只有在值和动态类型都为 nil 的情况下，接口值才为 nil。Foo() 函数返回的 err 变量，值为 nil、动态类型为 *os.PathError，与 nil（值为 nil，动态类型为 nil）显然是不相等。我们可以打印下变量 err 的详情：
//
//fmt.Printf("%#v\n",err)   // (*os.PathError)(nil)
//
//一个更合适的解决办法：
//
//func Foo() (err error) {
//	// …
//	return
//}
//
//func main() {
//	err := Foo()
//	fmt.Println(err)
//	fmt.Println(err == nil)
//}