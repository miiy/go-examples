//1.下面的代码有什么问题，请说明？
package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

//参考答案及解析：defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。
//
//修复代码：
//
//func main() {
//	f, err := os.Open("file")
//	if err != nil {
//		return
//	}
//	defer f.Close()
//
//	b, err := ioutil.ReadAll(f)
//	println(string(b))
//}
