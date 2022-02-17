//1.下面代码有什么问题？
package main

type foo struct {
	bar int
}

func main() {
	var f foo
	f.bar, tmp := 1, 2
}

//参考答案及解析：编译错误：
//
//non-name f.bar on left side of :=
//
//:= 操作符不能用于结构体字段赋值。
