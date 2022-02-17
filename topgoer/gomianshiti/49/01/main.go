//1.下面代码输出什么？
package main

func main() {
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}

//参考答案及解析：default。ch 为 nil，读写都会阻塞。
