//1.下面这段代码输出什么？
package main

import "fmt"

func main() {
	m := map[int]string{0:"zero",1:"one"}
	for k,v := range m {
		fmt.Println(k,v)
	}
}

//参考答案及解析：
//
//0 zero
//1 one
//// 或者
//1 one
//0 zero
//
//map 的输出是无序的。
