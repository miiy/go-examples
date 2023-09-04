//1.下面的代码输出什么？
package main

import "fmt"

type Point struct { x, y int }

func main()  {
	s := []Point{
		{1, 2},
		{3, 4},
	}
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(s)
}

//参考答案及解析：输出 [{1 2} {3 4}]。知识点：for range 循环。range 循环的时候，获取到的元素值是副本，就比如这里的 p。修复代码示例：
//
//type Point struct{ x, y int }
//
//func main()  {
//	s := []*Point{
//		&Point{1, 2},
//		&Point{3, 4},
//	}
//	for _, p := range s {
//		p.x, p.y = p.y, p.x
//	}
//	fmt.Println(*s[0])
//	fmt.Println(*s[1])
//}
