package main

import (
	"fmt"

	"github.com/miiy/go-examples/packages/stringer/demo2/enum"
)

func main() {
	fmt.Print(enum.ERR_CODE_OK)
	fmt.Print(enum.ERR_CODE_TIMEOUT)
	fmt.Println(int(enum.ERR_CODE_TIMEOUT))

}
