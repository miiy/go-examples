//1.下面代码能编译通过吗？请简要说明。
package main

import "fmt"

func main()  {
	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)
}

//参考答案及解析：能编译通过，输出 [1 2 3 0 1 2]。for 循环开始的时候，终止条件只会计算一次。
