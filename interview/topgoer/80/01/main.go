package main
//1.定义一个包内全局字符串变量，下面语法正确的是？
//A. var str string
//B. str := “”
//C. str = “”
//D. var str = “”
//
//参考答案及解析：AD。全局变量要定义在函数之外，而在函数之外定义的变量只能用 var 定义。短变量声明 := 只能用于函数之内。