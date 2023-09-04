//1.下面代码输出什么？
package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

//参考答案及解析：
//
//0xc000018058 2
//0xc000018058 2
//
//知识点：闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）。
