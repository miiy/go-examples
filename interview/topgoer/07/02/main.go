//2.下面这段代码能否编译通过？如果可以，输出什么？
package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}

//参考答案及解析：编译通过，输出：0 2 zz zz 5。知识点：iota 的使用。给大家贴篇文章，讲的很详细
//https://www.cnblogs.com/zsy/p/5370052.html
