//1.f1()、f2()、f3() 函数分别返回什么？
package main

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//答案是 29 29 28
//首先定义的局部变量person类型是一个指针
//其次defer是先进后出结构，故defer执行顺序为3 2 1
//3是匿名函数使用外部对象，而对象是指针，又在defer中，执行优先级为最低，故最外层代码修改以后，defer则会使用修改后的对象，故29
//2是函数传递参数进去，因为是指针，同3
//1看是一条语句，其实可以写成defer func(age int){fmt.Println(age)}(person.age) 因为传递的是执行到此初始person对象的age值，是而在此时age为28因为是值类型传递，所以输出为28
