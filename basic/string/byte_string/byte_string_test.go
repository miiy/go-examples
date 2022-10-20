package byte_string

import (
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	b1 := []byte("abc")
	for i := 0; i < b.N; i++ {
		BytesToString(b1)
	}
}

func BenchmarkBytesToString1(b *testing.B) {
	b1 := []byte("abc")
	for i := 0; i < b.N; i++ {
		BytesToString1(b1)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes("abc")
	}
}

func BenchmarkStringToBytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes1("abc")
	}
}
