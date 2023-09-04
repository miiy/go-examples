package main

import "fmt"
import "github.com/miiy/go-examples/pkg_stringer/2/enum"

func main() {
	fmt.Print(enum.ERR_CODE_OK)
	fmt.Print(enum.ERR_CODE_TIMEOUT)
	fmt.Println(int(enum.ERR_CODE_TIMEOUT))

}
