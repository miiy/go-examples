//2.下面代码有什么不规范的地方吗？
package main

import "fmt"

func main() {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}
}

//参考答案及解析：检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问 map 时返回的第二个值。
//
//修复代码如下：
//
//func main() {
//	x := map[string]string{"one":"a","two":"","three":"c"}
//	if _,ok := x["two"]; !ok {
//		fmt.Println("no entry")
//	}
//}
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
