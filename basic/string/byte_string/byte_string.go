package byte_string

import (
	"unsafe"
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func BytesToString1(b []byte) string {
	return string(b)
}

func StringToBytes1(s string) []byte {
	return []byte(s)
}
