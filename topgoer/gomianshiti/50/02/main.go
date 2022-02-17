//2.下面代码输出什么？
package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch(i) {
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

//参考答案及解析：false true。Go 语言的 switch 语句虽然没有”break”，但如果 case 完成程序会默认 break，可以在 case 语句后面加上关键字 fallthrough，这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。或者，利用 case 可以匹配多个值的特性。
//
//修复代码：
//
//func main() {
//	isMatch := func(i int) bool {
//		switch(i) {
//		case 1:
//			fallthrough
//		case 2:
//			return true
//		}
//		return false
//	}
//	fmt.Println(isMatch(1))     // true
//	fmt.Println(isMatch(2))     // true
//	match := func(i int) bool {
//		switch(i) {
//		case 1,2:
//			return true
//		}
//		return false
//	}
//	fmt.Println(match(1))       // true
//	fmt.Println(match(2))       // true
//}
