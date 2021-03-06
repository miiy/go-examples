//1.下面的代码有什么问题？
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
}

//参考答案及解析：导入的包没有被使用。如果引入一个包，但是未使用其中如何函数、接口、结构体或变量的话，代码将编译失败。
//如果你真的需要引入包，可以使用下划线操作符，_，来作为这个包的名字，从而避免失败。下划线操作符用于引入，但不使用。
//我们还可以注释或者移除未使用的包。
//
//修复代码：
//
//import (
//	_ "fmt"
//	"log"
//	"time"
//)
//var _ = log.Println
//func main() {
//	_ = time.Now
//}
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
