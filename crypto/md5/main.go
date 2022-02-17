package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main()  {
	str := "123456"
	fmt.Printf("%x\n", md5.Sum([]byte(str)))

	h := md5.New()
	h.Write([]byte(str))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

//Output:
//e10adc3949ba59abbe56e057f20f883e
//e10adc3949ba59abbe56e057f20f883e
