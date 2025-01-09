package main

import (
	"fmt"
	"math"
)

func main()  {
	// builtin.go
	/*

	布尔型：true, false

	整数：
	int8:  -128 -> 127
	int16: -32768 -> 32767
	int32: -2,147,483,648 -> 2,147,483,647
	int64: -9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807

	无符号整数：
	uint8:  0 -> 255
	uint16: 0 -> 65,535 万
	uint32: 0 -> 4,294,967,295
	uint64: 0 -> 18,446,744,073,709,551,615


	int32 21亿
	uint32 42亿左右

	*/

	//int8: -128 -> 127
	//int16: -32768 -> 32767
	//int32: -2147483648 -> 2147483647
	//int64: -9223372036854775808 -> 9223372036854775807
	fmt.Printf("int8: %d -> %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16: %d -> %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32: %d -> %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64: %d -> %d\n", math.MinInt64, math.MaxInt64)

	//uint8: 0 -> 255
	//uint16: 0 -> 65535
	//uint32: 0 -> 4294967295
	fmt.Printf("uint8: %d -> %d\n", 0, math.MaxUint8)
	fmt.Printf("uint16: %d -> %d\n", 0, math.MaxUint16)
	fmt.Printf("uint32: %d -> %d\n", 0, math.MaxUint32)
	// fmt.Printf("uint64: %d -> %d\n", 0, math.MaxUint64)
}
