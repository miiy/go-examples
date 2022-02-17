//3.下面代码能编译通过吗？
package main

import "fmt"

type info struct {
	result int
}
func work() (int,error) {
	return 13,nil
}
func main() {
	var data info
	data.result, err := work()
	fmt.Printf("info: %+v\n",data)
}

//参考答案及解析：编译失败。
//
//non-name data.result on left side of :=
//
//不能使用短变量声明设置结构体字段值，修复代码：
//
//func main() {
//	var data info
//	var err error
//	data.result, err = work() //ok
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(data)
//}
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
