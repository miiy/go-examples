package main

import (
	"fmt"
	"runtime"
)

/*
func NumCPU() int

NumCPU returns the number of logical CPUs usable by the current process.

The set of available CPUs is checked by querying the operating system at process startup.
Changes to operating system CPU allocation after process startup are not reflected.

NumCPU返回当前进程可用的逻辑cpu的数量。

通过在进程启动时查询操作系统来检查可用的cpu集。
进程启动后对操作系统CPU分配的更改不会反映出来。

*/

/*
func GOMAXPROCS(n int) int

GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
If n < 1, it does not change the current setting.
The number of logical CPUs on the local machine can be queried with NumCPU.
This call will go away when the scheduler improves.
*/
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu:", cpuNum)

	// Go 1.5开始， Go的GOMAXPROCS默认值已经设置为 CPU的核数
	// runtime.GOMAXPROCS(ncpu)
}

// cpu: 16
