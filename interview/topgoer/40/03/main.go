//3.下面这段代码存在什么问题？
package main

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	s.Param["day"] = 2
}

//参考答案及解析：存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引。修复代码如下：
//
//func main() {
//	s := new(Show)
//	// 修复代码
//	p := make(Param)
//	p["day"] = 2
//	s.Param = &p
//	tmp := *s.Param
//	fmt.Println(tmp["day"])
//}