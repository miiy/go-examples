package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	demo1()
	demo2()
}

// Output:
// SGVsbG8sIOS4lueVjA==
// Hello, 世界
// SGVsbG8sIFdvcmxkIQ==
// "Hello, World!"

func demo1() {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Printf("%s\n", decoded)
}

func demo2() {
	data := []byte("Hello, World!")
	edst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(edst, data)

	str := string(edst)
	fmt.Println(string(edst))

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	dst = dst[:n]
	fmt.Printf("%q\n", dst)
}
