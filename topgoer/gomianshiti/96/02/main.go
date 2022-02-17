//2.下面的代码有什么隐患？
package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

func main()  {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])
}

//参考答案及解析：get() 函数返回的切片与原切片公用底层数组，如果在调用函数里面（这里是 main() 函数）修改返回的切片，将会影响到原切片。为了避免掉入陷阱，可以如下修改：
//
//func get() []byte {
//	raw := make([]byte, 10000)
//	fmt.Println(len(raw), cap(raw), &raw[0])
//	res := make([]byte, 3)
//	copy(res, raw[:3])
//	return res
//}
//
//func main() {
//	data := get()
//	fmt.Println(len(data), cap(data), &data[0])
//}
