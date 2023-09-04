//3.下面代码能否编译通过？如果通过，输出什么？
package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {

	if _, exist := m[id]; exist {
		return "exist", true
	}
	return nil, false
}

func main() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}

//参考答案及解析：不能通过编译。知识点：函数返回值类型。nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错:cannot use nil as type string in return argument.
