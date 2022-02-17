//2.下面代码输出什么，为什么？
package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}

//参考答案及解析：recover:1。知识点：panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。
