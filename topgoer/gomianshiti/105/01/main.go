// 1.下面代码输出什么？请简要说明
package main

var c = make(chan int)
var a int

func f()  {
	a = 1
	<-c
}

func main()  {
	go f()
	c <- 0
	print(a)
}

//A. 不能编译
//B. 输出 1
//C. 输出 0
//D. panic
//
//参考答案及解析：B。能正确输出，不过主协程会阻塞 f() 函数的执行。