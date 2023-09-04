//2.下面代码输出什么？
package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}

//参考答案及解析：100 110。知识点：闭包引用相同变量。
