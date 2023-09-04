//1.下面哪一行代码会 panic，请说明原因？
package main

func main() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y
}

//参考答案及解析：第 8 行。因为两个比较值的动态类型为同一个不可比较类型。
