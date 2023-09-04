//1.下面这段代码输出什么？
package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}

//参考答案及解析：7。根据 5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！ 提到的“三步拆解法”，第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。
//
//此题引自知识星球《Go项目实战》。