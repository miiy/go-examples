//3.下面这段代码输出什么？
package main

import "fmt"

func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

//A. runtime panic
//B. 0
//C. compilation error
//
//参考答案及解析：B。删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。