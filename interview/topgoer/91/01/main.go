package main
//1.下面两段代码能否编译通过？请简要说明。
//第一段：
//
//func f() {}
//func f() {}
//func main() {}
//
//第二段：
//
//func init(){}
//func init(){}
//func main() {}
//
//参考答案及解析：第二段代码能通过编译。除 init() 函数之外，一个包内不允许有其他同名函数。
