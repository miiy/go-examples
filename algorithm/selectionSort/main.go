package selectionSort

func selectionSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
