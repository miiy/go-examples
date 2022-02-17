//2.下面代码能编译通过吗？可以的话，输出什么？
package main

func alwaysFalse() bool {
	return false
}

func main() {
	switch alwaysFalse()
	{
	case true:
		println(true)
	case false:
		println(false)
	}
}
//
//参考答案及解析：可以编译通过，输出：true。知识点：Go 代码断行规则。
//
//详情请查看：https://gfw.go101.org/article/line-break-rules.html
