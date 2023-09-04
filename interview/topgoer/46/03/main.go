//3.下面代码有什么问题？
package main

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
}

//参考答案及解析：将 nil 分配给 string 类型的变量。这是个大多数新手会犯的错误。修复代码：
//
//func main() {
//	var x string //defaults to "" (zero value)
//
//	if x == "" {
//		x = "default"
//	}
//}
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
