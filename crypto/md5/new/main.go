package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main()  {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!") // h.Write([]byte("And Leon's getting laaarger!"))
	fmt.Printf("%x", h.Sum(nil))
}

//Output:
//e2c569be17396eca2a2e3c11578123ed
