//2.下面代码输出什么？
package main

import "fmt"

func main()  {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}

//参考答案及解析：输出 []。知识点：拷贝切片。copy(dst, src) 函数返回 len(dst)、len(src) 之间的最小值。如果想要将 src 完全拷贝至 dst，必须给 dst 分配足够的内存空间。
//
//修复代码：
//
//func main() {
//	var src, dst []int
//	src = []int{1, 2, 3}
//	dst = make([]int, len(src))
//	n := copy(dst, src)
//	fmt.Println(n,dst)
//}
//
//或者直接使用 append()
//
//func main() {
//	var src, dst []int
//	src = []int{1, 2, 3}
//	dst = append(dst, src...)
//	fmt.Println("dst:", dst)
//}
//
//详情请参考《非懂不可的Slice（二）》https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA