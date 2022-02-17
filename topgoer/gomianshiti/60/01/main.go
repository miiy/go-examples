//1.下面哪一行代码会 panic，请说明原因？
package main

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	_ = s.foo
	s.foo()
	_ = s.bar
}

//参考答案及解析：第 19 行，因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。
//
//可以使用下面代码输出 s：
//
//func main() {
//	s := S{}
//	fmt.Printf("%#v",s)   // 输出：main.S{T:(*main.T)(nil)}
//}
//引自：《Go语言101》
