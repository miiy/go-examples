package main

import (
	"fmt"
)

// iota 是go语言中的常量计算器
func main()  {

	// 只能在常量表达式中使用
	// fmt.Println(iota) // undefined: iota

	const n = iota
	fmt.Println(n)

	const (
		a = iota // a = 0
		b        // b = 1
		c        // c = 2
		d = 5    // e = 5 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
		e        // e = 5
	)
	fmt.Println(a, b, c, d , e)

	const (
		_ = iota // 使用 _ 忽略不需要的 iota
		i        // i = 1
		j        // j = 2
		k        // k = 3
	)
	fmt.Println(i, j, k)


	type Stereotype int

	const (
		TypicalNoob Stereotype = iota // 0
		TypicalHipster                // 1   TypicalHipster = iota
		TypicalUnixWizard             // 2  TypicalUnixWizard = iota
		TypicalStartupFounder         // 3  TypicalStartupFounder = iota
	)
	fmt.Println(TypicalNoob)
	fmt.Println(TypicalHipster)
}
