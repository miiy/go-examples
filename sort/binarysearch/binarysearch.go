// https://ld246.com/article/1525348580697
package main

func main()  {
	
}

func binarySearch(values []int, start int, end int, key int) int {
	if start > end {
		return -1
	}

	mid := start + (end - start) / 2
	if values[mid] > key {
		return binarySearch()
	}
}