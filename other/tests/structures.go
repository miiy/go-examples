package main

import "fmt"

// 结构体定义：
// type identifier struct {
//     field1 type1
//     field2 type2
//     ...
// }

type struct1 struct {
	i1 int
	f1 float32
	str string
}

type T struct {
	a, b int
}

func main()  {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	var s T
	s.a = 5
	s.b = 8
	fmt.Println(s)

	var t *T
	t = new(T)
	fmt.Println(t)

}