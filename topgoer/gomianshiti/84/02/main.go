//2.下面代码输出什么？
package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	s := a[1:2]

	s[0] = 11
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}

//参考答案及解析：
//
//输出：
//
//[0 11 12]
//[21 12 13]
//
//详情请参考《非懂不可的Slice（二）》https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA
