//1.下面代码有什么问题？
package main

func main() {
	m := make(map[string]int,2)
	cap(m)
}

//参考答案及解析：问题：使用 cap() 获取 map 的容量。1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html