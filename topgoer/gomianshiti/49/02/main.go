//2.下面这段代码输出什么？
package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
         "name":"seekload"
     }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}

//参考答案及解析：输出 {}。知识点：结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体。修复代码：
//
//type People struct {
//	Name string `json:"name"`
//}
