//1.下面这段代码能通过编译吗？请简要说明。
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

//参考答案及解析：能通过编译。
//上面的代码可以理解成：
//
//func main() {
//	m := make(map[string]int)
//	v := m["foo"]
//	v++
//	m["foo"] = v
//	fmt.Println(m["foo"])
//}
