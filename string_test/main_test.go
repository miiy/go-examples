package main

import (
	"testing"
)

func benchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test1()
	}
}

func benchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test2()
	}
}
