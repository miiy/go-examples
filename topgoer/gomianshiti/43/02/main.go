//2.下面代码输出什么？
package main

func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
}

//A. true true false
//B. false true true
//C. true true true
//D. false true false
//
//参考答案及解析：D。知识点：类型断言。类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败
