// 2.下面的代码有几处问题？请详细说明。
package main

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}
func getT() T {
	return T{}
}

func main() {
	getT().Set(1)
}

//参考答案及解析：有两处问题：
//
//1.直接返回的 T{} 不可寻址；
//2.不可寻址的结构体不能调用带结构体指针接收者的方法；
//
//修复代码：
//
//type T struct {
//	n int
//}
//
//func (t *T) Set(n int) {
//	t.n = n
//}
//
//func getT() T {
//	return T{}
//}
//
//func main() {
//	t := getT()
//	t.Set(2)
//	fmt.Println(t.n)
//}
