package main
//1.关于map，下面说法正确的是？
//A. map 反序列化时 json.unmarshal() 的入参必须为map的地址；
//B. 在函数调用中传递 map，则子函数中对 map 元素的增加不会导致父函数中 map 的修改；
//C. 在函数调用中传递 map，则子函数中对 map 元素的修改不会导致父函数中 map 的修改；
//D. 不能使用内置函数 delete() 删除 map 的元素；
//
//参考答案及解析：A。知识点：map 的使用。可以查看《Go Map》
