//1.n 是秒数，下面代码输出什么？
package main

import "fmt"

func main() {
	n := 43210
	fmt.Println(n/60*60, " hours and ", n%60*60, " seconds")
}

//参考答案及解析：43200 hours and 600 seconds。知识点：运算符优先级。算术运算符 *、/ 和 % 的优先级相同，从左向右结合。
//
//修复代码如下：
//
//func main() {
//	n := 43210
//	fmt.Println(n/(60*60), "hours and", n%(60*60), "seconds")
//}
