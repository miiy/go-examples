package main

import (
	"fmt"
	"unicode"
)

func main()  {
	// 是否为字母
	fmt.Println(unicode.IsLetter('1'))
	fmt.Println(unicode.IsLetter('a'))
	// 是否为数字
	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(unicode.IsDigit('a'))
	// 是否为空白符号
	fmt.Println(unicode.IsSpace('1'))
	fmt.Println(unicode.IsSpace(' '))
}

/*
false
true
true
false
false
true
*/