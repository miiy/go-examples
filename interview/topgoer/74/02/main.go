package main
//2.假设 x 已声明，y 未声明，下面 4 行代码哪些是正确的。错误的请说明原因？
//x, _ := f()  // 1
//x, _ = f()  // 2
//x, y := f()  // 3
//x, y = f()  // 4
//
//参考答案及解析：2、3正确。知识点：简短变量声明。使用简短变量声明有几个需要注意的地方：
//
//只能用于函数内部；
//短变量声明语句中至少要声明一个新的变量；