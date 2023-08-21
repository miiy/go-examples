package main

import "fmt"

func main() {
	fmt.Println(SumIntsOrFloatsMap(map[string]int64{"a": 1, "b": 2}))
	fmt.Println(SumIntsOrFloatsMap(map[string]float64{"a": 1.1, "b": 2}))

	//fmt.Println(SumIntsOrFloats(1, 2))

	fmt.Println(SumIntsOrFloats[int64](1, 2))

	fmt.Println(SumNumbers(1.1, 2.1))
}

func SumIntsOrFloatsMap[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumIntsOrFloats[V int64 | float64](x, y V) V {
	return x + y
}

type Number interface {
	int64 | float64
}

func SumNumbers[V Number](x, y V) V {
	return x + y
}
