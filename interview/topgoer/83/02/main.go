//2.下面的代码有什么问题，请说明。
package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{"one"}
	d1.print()

	var in printer = data{"two"}
	in.print()
}

//参考答案及解析：编译报错。
//
//cannot use data literal (type data) as type printer in assignment:
//data does not implement printer (print method has pointer receiver)
//
//结构体类型 data 没有实现接口 printer。知识点：接口。
