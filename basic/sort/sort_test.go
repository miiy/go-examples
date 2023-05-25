package main

import (
	"fmt"
	"sort"
	"testing"
)

// go 默认提供了3个

func TestFloat64Sort(t *testing.T) {
	arr := []float64{0.1, 0.5, 0.8, 0.4, 0.2}
	sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
	fmt.Println(arr) // [0.8 0.5 0.4 0.2 0.1]
}
