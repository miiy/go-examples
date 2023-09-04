//2.下面哪一行代码会 panic，请说明原因？
package main

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	for range p {
		_ = len(p)
	}
	var s []int // nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
}

//参考答案及解析：第 12 行。因为左侧的 s[0] 中的 s 为 nil。
//
//引自：《Go语言101》
