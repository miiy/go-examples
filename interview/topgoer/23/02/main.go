// 2.下面选项正确的是？
package main

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}

//A. 1 2
//B. compilation error
//
//参考答案及解析：A。知识点：代码块和变量作用域。
